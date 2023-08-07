package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Concat concatenates two arrays into a single array.
// The function accepts two arrays of any type (interface{}) and returns a single concatenated array.
func Concat(array, extend interface{}) (interface{}, error) {
	arrValue, extendValue := reflect.ValueOf(array), reflect.ValueOf(extend)

	if (arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array) ||
		(extendValue.Kind() != reflect.Slice && extendValue.Kind() != reflect.Array) {
		return nil, constants.ErrNotSlice
	}

	if arrValue.Len() == 0 {
		return extend, nil
	}

	if extendValue.Len() == 0 {
		return array, nil
	}

	result, kind := reflect.MakeSlice(arrValue.Type(), 0, 0), arrValue.Index(0).Kind()
	for i := 0; i < arrValue.Len(); i++ {
		result = reflect.Append(result, arrValue.Index(i))
	}

	for i := 0; i < extendValue.Len(); i++ {
		element := extendValue.Index(i)
		if kind != element.Kind() && kind != reflect.Interface {
			return nil, constants.ErrIncompatible
		}
		result = reflect.Append(result, element)
	}

	return result.Interface(), nil
}
