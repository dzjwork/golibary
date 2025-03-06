package assert

import (
	"fmt"
	"reflect"
)

// 断言集合中元素是否是有序的
func isOrdered(t TestingT, object interface{}, allowedComparesResults []compareResult, failMessage string, msgAndArgs ...interface{}) bool {
	// 获取对象类型
	objKind := reflect.TypeOf(object).Kind()

	// 不是分片和数组直接退出
	if objKind != reflect.Slice && objKind != reflect.Array {
		return false
	}

	objValue := reflect.ValueOf(object)
	objLen := objValue.Len()

	// 空集合或只有一个元素的集合就是有序的
	if objLen <= 1 {
		return true
	}

	value := objValue.Index(0)
	valueInterface := value.Interface()
	firstValueKind := value.Kind()

	for i := 1; i < objLen; i++ {
		prevValue := value
		prevValueInterface := valueInterface

		value = objValue.Index(i)
		valueInterface = value.Interface()

		compareResult, isComparable := compare(prevValueInterface, valueInterface, firstValueKind)

		if !isComparable {
			return Fail(t, fmt.Sprintf("Can not compare type \"%s\" and \"%s\"", reflect.TypeOf(value), reflect.TypeOf(prevValue)), msgAndArgs...)
		}

		if !containsValue(allowedComparesResults, compareResult) {
			return Fail(t, fmt.Sprintf(failMessage, prevValue, value), msgAndArgs...)
		}
	}

	return true
}

// 断言集合中元素是否是升序
func IsIncreasing(t TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	return isOrdered(t, object, []compareResult{compareLess}, "\"%v\" is not less than \"%v\"", msgAndArgs...)
}

// 断言集合中元素是否是降序（可以相等）
func IsNonIncreasing(t TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	return isOrdered(t, object, []compareResult{compareEqual, compareGreater}, "\"%v\" is not greater than or equal to \"%v\"", msgAndArgs...)
}

// 断言集合中元素是否是降序
func IsDecreasing(t TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	return isOrdered(t, object, []compareResult{compareGreater}, "\"%v\" is not greater than \"%v\"", msgAndArgs...)
}

// 断言集合中元素是否是升序（可以相等）
func IsNonDecreasing(t TestingT, object interface{}, msgAndArgs ...interface{}) bool {
	return isOrdered(t, object, []compareResult{compareLess, compareEqual}, "\"%v\" is not less than or equal to \"%v\"", msgAndArgs...)
}
