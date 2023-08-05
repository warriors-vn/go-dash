package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// remove elements from the input array based on a given predicate function.
// It takes an array-like data structure and a predicate function that determines
// whether an element should be removed.
// The function returns the modified array and an error if any occurs.
func remove(array interface{}, predicate interface{}) (interface{}, error) {
	arrValue, predicateValue := reflect.ValueOf(array), reflect.ValueOf(predicate)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if predicateValue.Kind() != reflect.Func {
		return nil, constants.ErrNotFunction
	}

	numParams := predicateValue.Type().NumIn()
	if numParams != 1 {
		return nil, constants.ErrNotSupport
	}

	kind, result := predicateValue.Type().In(0).Kind(), reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := 0; i < arrValue.Len(); i++ {
		element := arrValue.Index(i)
		if element.Kind() != kind {
			return nil, constants.ErrIncompatible
		}

		res := predicateValue.Call([]reflect.Value{reflect.ValueOf(element.Interface())})
		if len(res) > 0 && res[0].Interface().(bool) {
			result = reflect.Append(result, element)
		}
	}

	return result.Interface(), nil
}
