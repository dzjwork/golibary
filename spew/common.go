/*
 * Copyright (c) 2013-2016 Dave Collins <dave@davec.name>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package spew

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
)

var (
	// PANIC=
	panicBytes = []byte("(PANIC=")
	// +
	plusBytes = []byte("+")
	// i
	iBytes = []byte("i")
	// true
	trueBytes = []byte("true")
	// false
	falseBytes = []byte("false")
	// (interface {})
	interfaceBytes = []byte("(interface {})")
	// ,\n
	commaNewlineBytes = []byte(",\n")
	// \n
	newlineBytes = []byte("\n")
	// {
	openBraceBytes = []byte("{")
	// {\n
	openBraceNewlineBytes = []byte("{\n")
	// }
	closeBraceBytes = []byte("}")
	// *
	asteriskBytes = []byte("*")
	// :
	colonBytes = []byte(":")
	// ：
	colonSpaceBytes = []byte(": ")
	// (
	openParenBytes = []byte("(")
	// )
	closeParenBytes = []byte(")")
	// 空格
	spaceBytes = []byte(" ")
	// ->
	pointerChainBytes = []byte("->")
	// <nil>
	nilAngleBytes = []byte("<nil>")
	// <max depth reached>\n
	maxNewlineBytes = []byte("<max depth reached>\n")
	// <max>
	maxShortBytes = []byte("<max>")
	// <already shown>
	circularBytes = []byte("<already shown>")
	// <shown>
	circularShortBytes = []byte("<shown>")
	// <invalid>
	invalidAngleBytes = []byte("<invalid>")
	// [
	openBracketBytes = []byte("[")
	// ]
	closeBracketBytes = []byte("]")
	// %
	percentBytes = []byte("%")
	// .
	precisionBytes = []byte(".")
	// <
	openAngleBytes = []byte("<")
	// >
	closeAngleBytes = []byte(">")
	// map[
	openMapBytes = []byte("map[")
	// ]
	closeMapBytes = []byte("]")
	// len=
	lenEqualsBytes = []byte("len=")
	// cap=
	capEqualsBytes = []byte("cap=")
)

// 十六进制中可以出现的字符
var hexDigits = "0123456789abcdef"

// 捕获到panic错误将其输出（必须在defer中使用）
func catchPanic(w io.Writer, v reflect.Value) {
	if err := recover(); err != nil {
		w.Write(panicBytes)
		fmt.Fprintf(w, "%v", err)
		w.Write(closeParenBytes)
	}
}

// handleMethods attempts to call the Error and String methods on the underlying
// type the passed reflect.Value represents and outputes the result to Writer w.
//
// It handles panics in any called methods by catching and displaying the error
// as the formatted value.
func handleMethods(cs *ConfigState, w io.Writer, v reflect.Value) (handled bool) {
	// We need an interface to check if the type implements the error or
	// Stringer interface.  However, the reflect package won't give us an
	// interface on certain things like unexported struct fields in order
	// to enforce visibility rules.  We use unsafe, when it's available,
	// to bypass these restrictions since this package does not mutate the
	// values.
	if !v.CanInterface() {
		if UnsafeDisabled {
			return false
		}

		v = unsafeReflectValue(v)
	}

	// Choose whether or not to do error and Stringer interface lookups against
	// the base type or a pointer to the base type depending on settings.
	// Technically calling one of these methods with a pointer receiver can
	// mutate the value, however, types which choose to satisify an error or
	// Stringer interface with a pointer receiver should not be mutating their
	// state inside these interface methods.
	if !cs.DisablePointerMethods && !UnsafeDisabled && !v.CanAddr() {
		v = unsafeReflectValue(v)
	}
	if v.CanAddr() {
		v = v.Addr()
	}

	// Is it an error or Stringer?
	switch iface := v.Interface().(type) {
	case error:
		defer catchPanic(w, v)
		if cs.ContinueOnMethod {
			w.Write(openParenBytes)
			w.Write([]byte(iface.Error()))
			w.Write(closeParenBytes)
			w.Write(spaceBytes)
			return false
		}

		w.Write([]byte(iface.Error()))
		return true

	case fmt.Stringer:
		defer catchPanic(w, v)
		if cs.ContinueOnMethod {
			w.Write(openParenBytes)
			w.Write([]byte(iface.String()))
			w.Write(closeParenBytes)
			w.Write(spaceBytes)
			return false
		}
		w.Write([]byte(iface.String()))
		return true
	}
	return false
}

// 打印布尔值
func printBool(w io.Writer, val bool) {
	if val {
		w.Write(trueBytes)
	} else {
		w.Write(falseBytes)
	}
}

// 打印int值
func printInt(w io.Writer, val int64, base int) {
	w.Write([]byte(strconv.FormatInt(val, base)))
}

// 打印Uint值
func printUint(w io.Writer, val uint64, base int) {
	w.Write([]byte(strconv.FormatUint(val, base)))
}

// 打印Float值
func printFloat(w io.Writer, val float64, precision int) {
	w.Write([]byte(strconv.FormatFloat(val, 'g', -1, precision)))
}

// 打印复数
func printComplex(w io.Writer, c complex128, floatPrecision int) {
	r := real(c)
	w.Write(openParenBytes)
	w.Write([]byte(strconv.FormatFloat(r, 'g', -1, floatPrecision)))
	i := imag(c)
	if i >= 0 {
		w.Write(plusBytes)
	}
	w.Write([]byte(strconv.FormatFloat(i, 'g', -1, floatPrecision)))
	w.Write(iBytes)
	w.Write(closeParenBytes)
}

// 输出十六进制的指针地址
func printHexPtr(w io.Writer, p uintptr) {
	// <nil>
	num := uint64(p)
	if num == 0 {
		w.Write(nilAngleBytes)
		return
	}

	// 64位 用16个十六进制字符 + 0x，一共需要18字节
	buf := make([]byte, 18)

	// 从右向左构建指针地址
	base := uint64(16)
	i := len(buf) - 1
	for num >= base {
		buf[i] = hexDigits[num%base]
		num /= base
		i--
	}
	buf[i] = hexDigits[num]

	// 0x前缀
	i--
	buf[i] = 'x'
	i--
	buf[i] = '0'

	// 删除没有使用的前导字节
	buf = buf[i:]
	w.Write(buf)
}

// valuesSorter implements sort.Interface to allow a slice of reflect.Value
// elements to be sorted.
type valuesSorter struct {
	values  []reflect.Value
	strings []string // either nil or same len and values
	cs      *ConfigState
}

// newValuesSorter initializes a valuesSorter instance, which holds a set of
// surrogate keys on which the data should be sorted.  It uses flags in
// ConfigState to decide if and how to populate those surrogate keys.
func newValuesSorter(values []reflect.Value, cs *ConfigState) sort.Interface {
	vs := &valuesSorter{values: values, cs: cs}
	if canSortSimply(vs.values[0].Kind()) {
		return vs
	}
	if !cs.DisableMethods {
		vs.strings = make([]string, len(values))
		for i := range vs.values {
			b := bytes.Buffer{}
			if !handleMethods(cs, &b, vs.values[i]) {
				vs.strings = nil
				break
			}
			vs.strings[i] = b.String()
		}
	}
	if vs.strings == nil && cs.SpewKeys {
		vs.strings = make([]string, len(values))
		for i := range vs.values {
			vs.strings[i] = Sprintf("%#v", vs.values[i].Interface())
		}
	}
	return vs
}

// canSortSimply tests whether a reflect.Kind is a primitive that can be sorted
// directly, or whether it should be considered for sorting by surrogate keys
// (if the ConfigState allows it).
func canSortSimply(kind reflect.Kind) bool {
	// This switch parallels valueSortLess, except for the default case.
	switch kind {
	case reflect.Bool:
		return true
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return true
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return true
	case reflect.Float32, reflect.Float64:
		return true
	case reflect.String:
		return true
	case reflect.Uintptr:
		return true
	case reflect.Array:
		return true
	}
	return false
}

// Len returns the number of values in the slice.  It is part of the
// sort.Interface implementation.
func (s *valuesSorter) Len() int {
	return len(s.values)
}

// Swap swaps the values at the passed indices.  It is part of the
// sort.Interface implementation.
func (s *valuesSorter) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
	if s.strings != nil {
		s.strings[i], s.strings[j] = s.strings[j], s.strings[i]
	}
}

// valueSortLess returns whether the first value should sort before the second
// value.  It is used by valueSorter.Less as part of the sort.Interface
// implementation.
func valueSortLess(a, b reflect.Value) bool {
	switch a.Kind() {
	case reflect.Bool:
		return !a.Bool() && b.Bool()
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return a.Int() < b.Int()
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return a.Uint() < b.Uint()
	case reflect.Float32, reflect.Float64:
		return a.Float() < b.Float()
	case reflect.String:
		return a.String() < b.String()
	case reflect.Uintptr:
		return a.Uint() < b.Uint()
	case reflect.Array:
		// Compare the contents of both arrays.
		l := a.Len()
		for i := 0; i < l; i++ {
			av := a.Index(i)
			bv := b.Index(i)
			if av.Interface() == bv.Interface() {
				continue
			}
			return valueSortLess(av, bv)
		}
	}
	return a.String() < b.String()
}

// Less returns whether the value at index i should sort before the
// value at index j.  It is part of the sort.Interface implementation.
func (s *valuesSorter) Less(i, j int) bool {
	if s.strings == nil {
		return valueSortLess(s.values[i], s.values[j])
	}
	return s.strings[i] < s.strings[j]
}

// sortValues is a sort function that handles both native types and any type that
// can be converted to error or Stringer.  Other inputs are sorted according to
// their Value.String() value to ensure display stability.
func sortValues(values []reflect.Value, cs *ConfigState) {
	if len(values) == 0 {
		return
	}
	sort.Sort(newValuesSorter(values, cs))
}
