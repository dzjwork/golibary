package spew

import (
	"reflect"
	"unsafe"
)

const (
	// 是否不允许访问不安全的包
	UnsafeDisabled = false

	// 指针大小
	ptrSize = unsafe.Sizeof((*byte)(nil))
)

// 指针类型
type flag uintptr

var (
	// 表示反射的值字段（只读的指针）
	flagRO flag

	// 表示反射的地址
	flagAddr flag
)

// 反射类型的掩码，用于计算反射类型
const flagKindMask = flag(0x1f)

// 不同的go使用不同标志位
var okFlags = []struct {
	ro, addr flag
}{{
	// From Go 1.4 to 1.5
	ro:   1 << 5,
	addr: 1 << 7,
}, {
	// Up to Go tip.
	ro:   1<<5 | 1<<6,
	addr: 1 << 8,
}}

// 获取发哦flag字段的字节偏移
var flagValOffset = func() uintptr {
	// 获取字段
	field, ok := reflect.TypeOf(reflect.Value{}).FieldByName("flag")
	if !ok {
		panic("reflect.Value has no flag field")
	}
	return field.Offset
}()

// 获取到指向反射的flag字段的指针
func flagField(v *reflect.Value) *flag {
	return (*flag)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + flagValOffset))
}

// unsafeReflectValue converts the passed reflect.Value into a one that bypasses
// the typical safety restrictions preventing access to unaddressable and
// unexported data.  It works by digging the raw pointer to the underlying
// value out of the protected value and generating a new unprotected (unsafe)
// reflect.Value to it.
//
// This allows us to check for implementations of the Stringer and error
// interfaces to be used for pretty printing ordinarily unaddressable and
// inaccessible values such as unexported struct fields.
func unsafeReflectValue(v reflect.Value) reflect.Value {
	if !v.IsValid() || (v.CanInterface() && v.CanAddr()) {
		return v
	}
	flagFieldPtr := flagField(&v)
	*flagFieldPtr &^= flagRO
	*flagFieldPtr |= flagAddr
	return v
}

// Sanity checks against future reflect package changes
// to the type or semantics of the Value.flag field.
func init() {
	field, ok := reflect.TypeOf(reflect.Value{}).FieldByName("flag")
	if !ok {
		panic("reflect.Value has no flag field")
	}
	if field.Type.Kind() != reflect.TypeOf(flag(0)).Kind() {
		panic("reflect.Value flag field has changed kind")
	}
	type t0 int
	var t struct {
		A t0
		// t0 will have flagEmbedRO set.
		t0
		// a will have flagStickyRO set
		a t0
	}
	vA := reflect.ValueOf(t).FieldByName("A")
	va := reflect.ValueOf(t).FieldByName("a")
	vt0 := reflect.ValueOf(t).FieldByName("t0")

	// Infer flagRO from the difference between the flags
	// for the (otherwise identical) fields in t.
	flagPublic := *flagField(&vA)
	flagWithRO := *flagField(&va) | *flagField(&vt0)
	flagRO = flagPublic ^ flagWithRO

	// Infer flagAddr from the difference between a value
	// taken from a pointer and not.
	vPtrA := reflect.ValueOf(&t).Elem().FieldByName("A")
	flagNoPtr := *flagField(&vA)
	flagPtr := *flagField(&vPtrA)
	flagAddr = flagNoPtr ^ flagPtr

	// Check that the inferred flags tally with one of the known versions.
	for _, f := range okFlags {
		if flagRO == f.ro && flagAddr == f.addr {
			return
		}
	}
	panic("reflect.Value read-only flag has changed semantics")
}
