package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Remove elements from the input array based on a given predicate function.
// It takes an array-like data structure and a predicate function that determines
// whether an element should be removed.
// The function returns the modified array and an error if any occurs.
func Remove(array interface{}, predicate interface{}) (interface{}, interface{}, error) {
	arrValue, predicateValue := reflect.ValueOf(array), reflect.ValueOf(predicate)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, nil, constants.ErrNotSlice
	}

	if predicateValue.Kind() != reflect.Func {
		return nil, nil, constants.ErrNotFunction
	}

	numParams := predicateValue.Type().NumIn()
	if numParams != 1 {
		return nil, nil, constants.ErrNotSupport
	}

	kind, result, old := predicateValue.Type().In(0).Kind(), reflect.MakeSlice(arrValue.Type(), 0, 0), reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := 0; i < arrValue.Len(); i++ {
		element := arrValue.Index(i)
		if element.Kind() != kind {
			return nil, nil, constants.ErrIncompatible
		}

		res := predicateValue.Call([]reflect.Value{reflect.ValueOf(element.Interface())})
		if len(res) > 0 && res[0].Kind() == reflect.Bool {
			if res[0].Interface().(bool) {
				result = reflect.Append(result, element)
			} else {
				old = reflect.Append(old, element)
			}
		}
	}

	return result.Interface(), old.Interface(), nil
}
