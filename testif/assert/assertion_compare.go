package assert

import (
	"bytes"
	"fmt"
	"reflect"
	"time"
)

type compareResult int

type compareType = compareResult

const (
	// 表示obj1 > obj2
	compareLess compareResult = iota - 1
	// 表示两个对象一样
	compareEqual
	// 表示obj1 < obj2
	compareGreater
)

var (
	// 整数类型
	intType = reflect.TypeOf(int(1))
	// 8位整数类型
	int8Type = reflect.TypeOf(int8(1))
	// 16位整数类型
	int16Type = reflect.TypeOf(int16(1))
	// 32位整数类型
	int32Type = reflect.TypeOf(int32(1))
	// 64位整数类型
	int64Type = reflect.TypeOf(int64(1))

	// 无符号整数类型
	uintType = reflect.TypeOf(uint(1))
	// 无符号8位整数类型
	uint8Type = reflect.TypeOf(uint8(1))
	// 无符号16位整数类型
	uint16Type = reflect.TypeOf(uint16(1))
	// 无符号32位整数类型
	uint32Type = reflect.TypeOf(uint32(1))
	// 无符号64位整数类型
	uint64Type = reflect.TypeOf(uint64(1))

	// 无符号指针类型
	uintptrType = reflect.TypeOf(uintptr(1))

	// 32位浮点数类型
	float32Type = reflect.TypeOf(float32(1))
	// 64位浮点数类型
	float64Type = reflect.TypeOf(float64(1))

	// 字符串类型
	stringType = reflect.TypeOf("")

	// 时间类型
	timeType = reflect.TypeOf(time.Time{})
	// 字节类型
	bytesType = reflect.TypeOf([]byte{})
)

// 比较两个对象的大小
func compare(obj1, obj2 interface{}, kind reflect.Kind) (compareResult, bool) {
	obj1Value := reflect.ValueOf(obj1)
	obj2Value := reflect.ValueOf(obj2)

	switch kind {
	case reflect.Int:
		{
			intobj1, ok := obj1.(int)
			if !ok {
				intobj1 = obj1Value.Convert(intType).Interface().(int)
			}
			intobj2, ok := obj2.(int)
			if !ok {
				intobj2 = obj2Value.Convert(intType).Interface().(int)
			}
			if intobj1 > intobj2 {
				return compareGreater, true
			}
			if intobj1 == intobj2 {
				return compareEqual, true
			}
			if intobj1 < intobj2 {
				return compareLess, true
			}
		}
	case reflect.Int8:
		{
			int8obj1, ok := obj1.(int8)
			if !ok {
				int8obj1 = obj1Value.Convert(int8Type).Interface().(int8)
			}
			int8obj2, ok := obj2.(int8)
			if !ok {
				int8obj2 = obj2Value.Convert(int8Type).Interface().(int8)
			}
			if int8obj1 > int8obj2 {
				return compareGreater, true
			}
			if int8obj1 == int8obj2 {
				return compareEqual, true
			}
			if int8obj1 < int8obj2 {
				return compareLess, true
			}
		}
	case reflect.Int16:
		{
			int16obj1, ok := obj1.(int16)
			if !ok {
				int16obj1 = obj1Value.Convert(int16Type).Interface().(int16)
			}
			int16obj2, ok := obj2.(int16)
			if !ok {
				int16obj2 = obj2Value.Convert(int16Type).Interface().(int16)
			}
			if int16obj1 > int16obj2 {
				return compareGreater, true
			}
			if int16obj1 == int16obj2 {
				return compareEqual, true
			}
			if int16obj1 < int16obj2 {
				return compareLess, true
			}
		}
	case reflect.Int32:
		{
			int32obj1, ok := obj1.(int32)
			if !ok {
				int32obj1 = obj1Value.Convert(int32Type).Interface().(int32)
			}
			int32obj2, ok := obj2.(int32)
			if !ok {
				int32obj2 = obj2Value.Convert(int32Type).Interface().(int32)
			}
			if int32obj1 > int32obj2 {
				return compareGreater, true
			}
			if int32obj1 == int32obj2 {
				return compareEqual, true
			}
			if int32obj1 < int32obj2 {
				return compareLess, true
			}
		}
	case reflect.Int64:
		{
			int64obj1, ok := obj1.(int64)
			if !ok {
				int64obj1 = obj1Value.Convert(int64Type).Interface().(int64)
			}
			int64obj2, ok := obj2.(int64)
			if !ok {
				int64obj2 = obj2Value.Convert(int64Type).Interface().(int64)
			}
			if int64obj1 > int64obj2 {
				return compareGreater, true
			}
			if int64obj1 == int64obj2 {
				return compareEqual, true
			}
			if int64obj1 < int64obj2 {
				return compareLess, true
			}
		}
	case reflect.Uint:
		{
			uintobj1, ok := obj1.(uint)
			if !ok {
				uintobj1 = obj1Value.Convert(uintType).Interface().(uint)
			}
			uintobj2, ok := obj2.(uint)
			if !ok {
				uintobj2 = obj2Value.Convert(uintType).Interface().(uint)
			}
			if uintobj1 > uintobj2 {
				return compareGreater, true
			}
			if uintobj1 == uintobj2 {
				return compareEqual, true
			}
			if uintobj1 < uintobj2 {
				return compareLess, true
			}
		}
	case reflect.Uint8:
		{
			uint8obj1, ok := obj1.(uint8)
			if !ok {
				uint8obj1 = obj1Value.Convert(uint8Type).Interface().(uint8)
			}
			uint8obj2, ok := obj2.(uint8)
			if !ok {
				uint8obj2 = obj2Value.Convert(uint8Type).Interface().(uint8)
			}
			if uint8obj1 > uint8obj2 {
				return compareGreater, true
			}
			if uint8obj1 == uint8obj2 {
				return compareEqual, true
			}
			if uint8obj1 < uint8obj2 {
				return compareLess, true
			}
		}
	case reflect.Uint16:
		{
			uint16obj1, ok := obj1.(uint16)
			if !ok {
				uint16obj1 = obj1Value.Convert(uint16Type).Interface().(uint16)
			}
			uint16obj2, ok := obj2.(uint16)
			if !ok {
				uint16obj2 = obj2Value.Convert(uint16Type).Interface().(uint16)
			}
			if uint16obj1 > uint16obj2 {
				return compareGreater, true
			}
			if uint16obj1 == uint16obj2 {
				return compareEqual, true
			}
			if uint16obj1 < uint16obj2 {
				return compareLess, true
			}
		}
	case reflect.Uint32:
		{
			uint32obj1, ok := obj1.(uint32)
			if !ok {
				uint32obj1 = obj1Value.Convert(uint32Type).Interface().(uint32)
			}
			uint32obj2, ok := obj2.(uint32)
			if !ok {
				uint32obj2 = obj2Value.Convert(uint32Type).Interface().(uint32)
			}
			if uint32obj1 > uint32obj2 {
				return compareGreater, true
			}
			if uint32obj1 == uint32obj2 {
				return compareEqual, true
			}
			if uint32obj1 < uint32obj2 {
				return compareLess, true
			}
		}
	case reflect.Uint64:
		{
			uint64obj1, ok := obj1.(uint64)
			if !ok {
				uint64obj1 = obj1Value.Convert(uint64Type).Interface().(uint64)
			}
			uint64obj2, ok := obj2.(uint64)
			if !ok {
				uint64obj2 = obj2Value.Convert(uint64Type).Interface().(uint64)
			}
			if uint64obj1 > uint64obj2 {
				return compareGreater, true
			}
			if uint64obj1 == uint64obj2 {
				return compareEqual, true
			}
			if uint64obj1 < uint64obj2 {
				return compareLess, true
			}
		}
	case reflect.Float32:
		{
			float32obj1, ok := obj1.(float32)
			if !ok {
				float32obj1 = obj1Value.Convert(float32Type).Interface().(float32)
			}
			float32obj2, ok := obj2.(float32)
			if !ok {
				float32obj2 = obj2Value.Convert(float32Type).Interface().(float32)
			}
			if float32obj1 > float32obj2 {
				return compareGreater, true
			}
			if float32obj1 == float32obj2 {
				return compareEqual, true
			}
			if float32obj1 < float32obj2 {
				return compareLess, true
			}
		}
	case reflect.Float64:
		{
			float64obj1, ok := obj1.(float64)
			if !ok {
				float64obj1 = obj1Value.Convert(float64Type).Interface().(float64)
			}
			float64obj2, ok := obj2.(float64)
			if !ok {
				float64obj2 = obj2Value.Convert(float64Type).Interface().(float64)
			}
			if float64obj1 > float64obj2 {
				return compareGreater, true
			}
			if float64obj1 == float64obj2 {
				return compareEqual, true
			}
			if float64obj1 < float64obj2 {
				return compareLess, true
			}
		}
	case reflect.String:
		{
			stringobj1, ok := obj1.(string)
			if !ok {
				stringobj1 = obj1Value.Convert(stringType).Interface().(string)
			}
			stringobj2, ok := obj2.(string)
			if !ok {
				stringobj2 = obj2Value.Convert(stringType).Interface().(string)
			}
			if stringobj1 > stringobj2 {
				return compareGreater, true
			}
			if stringobj1 == stringobj2 {
				return compareEqual, true
			}
			if stringobj1 < stringobj2 {
				return compareLess, true
			}
		}
	// Check for known struct types we can check for compare results.
	case reflect.Struct:
		{
			// All structs enter here. We're not interested in most types.
			if !obj1Value.CanConvert(timeType) {
				break
			}

			// time.Time can be compared!
			timeObj1, ok := obj1.(time.Time)
			if !ok {
				timeObj1 = obj1Value.Convert(timeType).Interface().(time.Time)
			}

			timeObj2, ok := obj2.(time.Time)
			if !ok {
				timeObj2 = obj2Value.Convert(timeType).Interface().(time.Time)
			}

			if timeObj1.Before(timeObj2) {
				return compareLess, true
			}
			if timeObj1.Equal(timeObj2) {
				return compareEqual, true
			}
			return compareGreater, true
		}
	case reflect.Slice:
		{
			// We only care about the []byte type.
			if !obj1Value.CanConvert(bytesType) {
				break
			}

			// []byte can be compared!
			bytesObj1, ok := obj1.([]byte)
			if !ok {
				bytesObj1 = obj1Value.Convert(bytesType).Interface().([]byte)

			}
			bytesObj2, ok := obj2.([]byte)
			if !ok {
				bytesObj2 = obj2Value.Convert(bytesType).Interface().([]byte)
			}

			return compareResult(bytes.Compare(bytesObj1, bytesObj2)), true
		}
	case reflect.Uintptr:
		{
			uintptrObj1, ok := obj1.(uintptr)
			if !ok {
				uintptrObj1 = obj1Value.Convert(uintptrType).Interface().(uintptr)
			}
			uintptrObj2, ok := obj2.(uintptr)
			if !ok {
				uintptrObj2 = obj2Value.Convert(uintptrType).Interface().(uintptr)
			}
			if uintptrObj1 > uintptrObj2 {
				return compareGreater, true
			}
			if uintptrObj1 == uintptrObj2 {
				return compareEqual, true
			}
			if uintptrObj1 < uintptrObj2 {
				return compareLess, true
			}
		}
	}

	return compareEqual, false
}

// 断言e1是否大于e2
func Greater(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return compareTwoValues(t, e1, e2, []compareResult{compareGreater}, "\"%v\" is not greater than \"%v\"", msgAndArgs...)
}

// 断言e1是否等于e2
func GreaterOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return compareTwoValues(t, e1, e2, []compareResult{compareGreater, compareEqual}, "\"%v\" is not greater than or equal to \"%v\"", msgAndArgs...)
}

// 断言e1是否小于e2
func Less(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return compareTwoValues(t, e1, e2, []compareResult{compareLess}, "\"%v\" is not less than \"%v\"", msgAndArgs...)
}

// 断言e1是否小于等于e2
func LessOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	return compareTwoValues(t, e1, e2, []compareResult{compareLess, compareEqual}, "\"%v\" is not less than or equal to \"%v\"", msgAndArgs...)
}

// 断言指定的值是否为整数
func Positive(t TestingT, e interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	// 获取0值
	zero := reflect.Zero(reflect.TypeOf(e))
	// 比较e是否小于0，如果是返回true，否则返回false
	return compareTwoValues(t, e, zero.Interface(), []compareResult{compareGreater}, "\"%v\" is not positive", msgAndArgs...)
}

// 判断指定的是否为负数
func Negative(t TestingT, e interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	// 获取类型的0值
	zero := reflect.Zero(reflect.TypeOf(e))
	// 比较e是否小于0，如果是返回true，否则返回false
	return compareTwoValues(t, e, zero.Interface(), []compareResult{compareLess}, "\"%v\" is not negative", msgAndArgs...)
}

// 比较两个值并查看值是否是期望的结果，如果是返回true，否则返回false
func compareTwoValues(t TestingT, e1 interface{}, e2 interface{}, allowedComparesResults []compareResult, failMessage string, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	e1Kind := reflect.ValueOf(e1).Kind()
	e2Kind := reflect.ValueOf(e2).Kind()
	// 类型不同无法比较
	if e1Kind != e2Kind {
		return Fail(t, "Elements should be the same type", msgAndArgs...)
	}

	// 比较两个对象大小，比较失败直接退出
	compareResult, isComparable := compare(e1, e2, e1Kind)
	if !isComparable {
		return Fail(t, fmt.Sprintf("Can not compare type \"%s\"", reflect.TypeOf(e1)), msgAndArgs...)
	}

	// 如果比较结果不在期望的结果集中直接退出
	if !containsValue(allowedComparesResults, compareResult) {
		return Fail(t, fmt.Sprintf(failMessage, e1, e2), msgAndArgs...)
	}

	return true
}

// 查看比较结果是否在期望的比较结果集中
func containsValue(values []compareResult, value compareResult) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}
