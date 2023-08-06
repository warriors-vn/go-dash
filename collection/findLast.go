package collection

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// findLast returns the last element from the input array that satisfies the predicate.
// It takes an array-like data structure and a predicate function that determines whether an element is a match.
// The predicate function should have the signature func(elementType) bool.
// The function returns the found element and a boolean indicating whether a match was found, or an error if any occurs.
func findLast(array interface{}, predicate interface{}) (interface{}, error) {
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

	kind := predicateValue.Type().In(0).Kind()
	for i := arrValue.Len() - 1; i >= 0; i-- {
		element := arrValue.Index(i)
		if element.Kind() != kind {
			return nil, constants.ErrIncompatible
		}

		res := predicateValue.Call([]reflect.Value{reflect.ValueOf(element.Interface())})
		if len(res) > 0 && res[0].Kind() == reflect.Bool && res[0].Interface().(bool) {
			return element.Interface(), nil
		}
	}

	return nil, constants.ErrNotFound
}
