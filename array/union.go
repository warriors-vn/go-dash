package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Union returns a new Slice that contains the unique elements from multiple input arrays.
// It takes variadic arguments representing array-like data structures and returns a new Slice
// containing unique elements from all input arrays combined.
// The function returns the new Slice and an error if any occurs.
func Union(arrays ...interface{}) (interface{}, error) {
	if arrays == nil {
		return nil, constants.ErrNotSupport
	}

	arrFirstValueOf := reflect.ValueOf(arrays[0])
	if arrFirstValueOf.Kind() != reflect.Slice && arrFirstValueOf.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if arrFirstValueOf.Len() == 0 {
		return nil, constants.ErrNotSupport
	}

	mapValue, result, firstKind := make(map[interface{}]bool), reflect.MakeSlice(arrFirstValueOf.Type(), 0, 0), arrFirstValueOf.Index(0).Kind()
	for i := 0; i < len(arrays); i++ {
		arr := reflect.ValueOf(arrays[i])
		if arr.Kind() != reflect.Slice && arr.Kind() != reflect.Array {
			continue
		}

		for j := 0; j < arr.Len(); j++ {
			element := arr.Index(j)
			if !mapValue[element.Interface()] {
				mapValue[element.Interface()] = true
				if firstKind != element.Kind() {
					return nil, constants.ErrIncompatible
				}
				result = reflect.Append(result, element)
			}
		}
	}

	return result.Interface(), nil
}
