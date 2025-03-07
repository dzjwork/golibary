package spew

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// fmt包支持的所有字符标志
const supportedFlags = "0-+# "

// 该结构实现了fmt，包含了格式化的状态
type formatState struct {
	value          interface{} // 需要被格式打印的内容
	fs             fmt.State   // 可以打印格式化的输出、设置标志位、查看宽度、精度标志位
	depth          int
	pointers       map[uintptr]int
	ignoreNextType bool
	cs             *ConfigState
}

// buildDefaultFormat recreates the original format string without precision
// and width information to pass in to fmt.Sprintf in the case of an
// unrecognized type.  Unless new types are added to the language, this
// function won't ever be called.
func (f *formatState) buildDefaultFormat() (format string) {
	buf := bytes.NewBuffer(percentBytes)

	for _, flag := range supportedFlags {
		if f.fs.Flag(int(flag)) {
			buf.WriteRune(flag)
		}
	}

	buf.WriteRune('v')

	format = buf.String()
	return format
}

// 创建包含精度、宽度、支持的占位符字符串
func (f *formatState) constructOrigFormat(verb rune) (format string) {
	buf := bytes.NewBuffer(percentBytes)

	// 查看State是否支持对应的占位符，如果支持则加到buf中
	for _, flag := range supportedFlags {
		if f.fs.Flag(int(flag)) {
			buf.WriteRune(flag)
		}
	}

	// 获取到State的宽度
	if width, ok := f.fs.Width(); ok {
		buf.WriteString(strconv.Itoa(width))
	}

	// 获取到State的精度
	if precision, ok := f.fs.Precision(); ok {
		buf.Write(precisionBytes)
		buf.WriteString(strconv.Itoa(precision))
	}

	buf.WriteRune(verb)

	format = buf.String()
	return format
}

// unpackValue returns values inside of non-nil interfaces when possible and
// ensures that types for values which have been unpacked from an interface
// are displayed when the show types flag is also set.
// This is useful for data types like structs, arrays, slices, and maps which
// can contain varying types packed inside an interface.
func (f *formatState) unpackValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Interface {
		f.ignoreNextType = false
		if !v.IsNil() {
			v = v.Elem()
		}
	}
	return v
}

// 格式化指针类型的打印
func (f *formatState) formatPtr(v reflect.Value) {
	// 如果指针为nil则显示<nil>
	showTypes := f.fs.Flag('#')
	if v.IsNil() && (!showTypes || f.ignoreNextType) {
		f.fs.Write(nilAngleBytes)
		return
	}

	// Remove pointers at or below the current depth from map used to detect
	// circular refs.
	for k, depth := range f.pointers {
		if depth >= f.depth {
			delete(f.pointers, k)
		}
	}

	// Keep list of all dereferenced pointers to possibly show later.
	pointerChain := make([]uintptr, 0)

	// Figure out how many levels of indirection there are by derferencing
	// pointers and unpacking interfaces down the chain while detecting circular
	// references.
	nilFound := false
	cycleFound := false
	indirects := 0
	ve := v
	for ve.Kind() == reflect.Ptr {
		if ve.IsNil() {
			nilFound = true
			break
		}
		indirects++
		addr := ve.Pointer()
		pointerChain = append(pointerChain, addr)
		if pd, ok := f.pointers[addr]; ok && pd < f.depth {
			cycleFound = true
			indirects--
			break
		}
		f.pointers[addr] = f.depth

		ve = ve.Elem()
		if ve.Kind() == reflect.Interface {
			if ve.IsNil() {
				nilFound = true
				break
			}
			ve = ve.Elem()
		}
	}

	// Display type or indirection level depending on flags.
	if showTypes && !f.ignoreNextType {
		f.fs.Write(openParenBytes)
		f.fs.Write(bytes.Repeat(asteriskBytes, indirects))
		f.fs.Write([]byte(ve.Type().String()))
		f.fs.Write(closeParenBytes)
	} else {
		if nilFound || cycleFound {
			indirects += strings.Count(ve.Type().String(), "*")
		}
		f.fs.Write(openAngleBytes)
		f.fs.Write([]byte(strings.Repeat("*", indirects)))
		f.fs.Write(closeAngleBytes)
	}

	// Display pointer information depending on flags.
	if f.fs.Flag('+') && (len(pointerChain) > 0) {
		f.fs.Write(openParenBytes)
		for i, addr := range pointerChain {
			if i > 0 {
				f.fs.Write(pointerChainBytes)
			}
			printHexPtr(f.fs, addr)
		}
		f.fs.Write(closeParenBytes)
	}

	// Display dereferenced value.
	switch {
	case nilFound:
		f.fs.Write(nilAngleBytes)

	case cycleFound:
		f.fs.Write(circularShortBytes)

	default:
		f.ignoreNextType = true
		f.format(ve)
	}
}

// format is the main workhorse for providing the Formatter interface.  It
// uses the passed reflect value to figure out what kind of object we are
// dealing with and formats it appropriately.  It is a recursive function,
// however circular data structures are detected and handled properly.
func (f *formatState) format(v reflect.Value) {
	// 如果v没有正确的被初始化直接输出<invalid>
	kind := v.Kind()
	if kind == reflect.Invalid {
		f.fs.Write(invalidAngleBytes)
		return
	}

	// 指针类型打印
	if kind == reflect.Ptr {
		f.formatPtr(v)
		return
	}

	// Print type information unless already handled elsewhere.
	if !f.ignoreNextType && f.fs.Flag('#') {
		f.fs.Write(openParenBytes)
		f.fs.Write([]byte(v.Type().String()))
		f.fs.Write(closeParenBytes)
	}
	f.ignoreNextType = false

	// Call Stringer/error interfaces if they exist and the handle methods
	// flag is enabled.
	if !f.cs.DisableMethods {
		if (kind != reflect.Invalid) && (kind != reflect.Interface) {
			if handled := handleMethods(f.cs, f.fs, v); handled {
				return
			}
		}
	}

	switch kind {
	case reflect.Invalid:
		// Do nothing.  We should never get here since invalid has already
		// been handled above.

	case reflect.Bool:
		printBool(f.fs, v.Bool())

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		printInt(f.fs, v.Int(), 10)

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		printUint(f.fs, v.Uint(), 10)

	case reflect.Float32:
		printFloat(f.fs, v.Float(), 32)

	case reflect.Float64:
		printFloat(f.fs, v.Float(), 64)

	case reflect.Complex64:
		printComplex(f.fs, v.Complex(), 32)

	case reflect.Complex128:
		printComplex(f.fs, v.Complex(), 64)

	case reflect.Slice:
		if v.IsNil() {
			f.fs.Write(nilAngleBytes)
			break
		}
		fallthrough

	case reflect.Array:
		f.fs.Write(openBracketBytes)
		f.depth++
		if (f.cs.MaxDepth != 0) && (f.depth > f.cs.MaxDepth) {
			f.fs.Write(maxShortBytes)
		} else {
			numEntries := v.Len()
			for i := 0; i < numEntries; i++ {
				if i > 0 {
					f.fs.Write(spaceBytes)
				}
				f.ignoreNextType = true
				f.format(f.unpackValue(v.Index(i)))
			}
		}
		f.depth--
		f.fs.Write(closeBracketBytes)

	case reflect.String:
		f.fs.Write([]byte(v.String()))

	case reflect.Interface:
		// The only time we should get here is for nil interfaces due to
		// unpackValue calls.
		if v.IsNil() {
			f.fs.Write(nilAngleBytes)
		}

	case reflect.Ptr:
		// Do nothing.  We should never get here since pointers have already
		// been handled above.

	case reflect.Map:
		// nil maps should be indicated as different than empty maps
		if v.IsNil() {
			f.fs.Write(nilAngleBytes)
			break
		}

		f.fs.Write(openMapBytes)
		f.depth++
		if (f.cs.MaxDepth != 0) && (f.depth > f.cs.MaxDepth) {
			f.fs.Write(maxShortBytes)
		} else {
			keys := v.MapKeys()
			if f.cs.SortKeys {
				sortValues(keys, f.cs)
			}
			for i, key := range keys {
				if i > 0 {
					f.fs.Write(spaceBytes)
				}
				f.ignoreNextType = true
				f.format(f.unpackValue(key))
				f.fs.Write(colonBytes)
				f.ignoreNextType = true
				f.format(f.unpackValue(v.MapIndex(key)))
			}
		}
		f.depth--
		f.fs.Write(closeMapBytes)

	case reflect.Struct:
		numFields := v.NumField()
		f.fs.Write(openBraceBytes)
		f.depth++
		if (f.cs.MaxDepth != 0) && (f.depth > f.cs.MaxDepth) {
			f.fs.Write(maxShortBytes)
		} else {
			vt := v.Type()
			for i := 0; i < numFields; i++ {
				if i > 0 {
					f.fs.Write(spaceBytes)
				}
				vtf := vt.Field(i)
				if f.fs.Flag('+') || f.fs.Flag('#') {
					f.fs.Write([]byte(vtf.Name))
					f.fs.Write(colonBytes)
				}
				f.format(f.unpackValue(v.Field(i)))
			}
		}
		f.depth--
		f.fs.Write(closeBraceBytes)

	case reflect.Uintptr:
		printHexPtr(f.fs, uintptr(v.Uint()))

	case reflect.UnsafePointer, reflect.Chan, reflect.Func:
		printHexPtr(f.fs, v.Pointer())

	// There were not any other types at the time this code was written, but
	// fall back to letting the default fmt package handle it if any get added.
	default:
		format := f.buildDefaultFormat()
		if v.CanInterface() {
			fmt.Fprintf(f.fs, format, v.Interface())
		} else {
			fmt.Fprintf(f.fs, format, v.String())
		}
	}
}

// 实现了fmt.Formatter接口
func (f *formatState) Format(fs fmt.State, verb rune) {
	f.fs = fs

	// 对除了v占位符以外的占位符进行格式化打印
	if verb != 'v' {
		format := f.constructOrigFormat(verb)
		fmt.Fprintf(fs, format, f.value)
		return
	}

	// 如果打印的值是空，并且有#标识符则输出(interface {}) <nil>，如果没有#标识符输出<nil>
	if f.value == nil {
		if fs.Flag('#') {
			fs.Write(interfaceBytes)
		}
		fs.Write(nilAngleBytes)
		return
	}

	// 格式化输出对象的类型
	f.format(reflect.ValueOf(f.value))
}

// 创建一个新的格式化器
func newFormatter(cs *ConfigState, v interface{}) fmt.Formatter {
	fs := &formatState{value: v, cs: cs}
	fs.pointers = make(map[uintptr]int)
	return fs
}

// 创建一个新的格式化器
func NewFormatter(v interface{}) fmt.Formatter {
	return newFormatter(&Config, v)
}
