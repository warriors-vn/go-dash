package collection

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Every check if all elements in the input array satisfy the predicate function.
// It takes an array-like data structure and a predicate function that determines the condition.
// The predicate function should have the signature func(elementType) bool.
// The function returns true if all elements satisfy the condition, false otherwise, or an error if any occurs.
func Every(array, predicate interface{}) (bool, error) {
	arrValue, predicateValue := reflect.ValueOf(array), reflect.ValueOf(predicate)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return false, constants.ErrNotSlice
	}

	if predicateValue.Kind() != reflect.Func {
		return false, constants.ErrNotFunction
	}

	numParams := predicateValue.Type().NumIn()
	if numParams != 1 {
		return false, constants.ErrNotSupport
	}

	kind := predicateValue.Type().In(0).Kind()
	for i := 0; i < arrValue.Len(); i++ {
		element := arrValue.Index(i)
		if element.Kind() != kind {
			return false, constants.ErrIncompatible
		}

		res := predicateValue.Call([]reflect.Value{reflect.ValueOf(element.Interface())})
		if len(res) > 0 && res[0].Kind() != reflect.Bool {
			return false, constants.ErrNotSupport
		}

		if len(res) > 0 && res[0].Kind() == reflect.Bool && !res[0].Interface().(bool) {
			return false, nil
		}
	}

	return true, nil
}
