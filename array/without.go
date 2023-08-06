package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// without returns a new slice that excludes specified values from the input array.
// It takes an array-like data structure and a variable number of values to be excluded.
// The function returns the new slice without the specified values and an error if any occurs.
func without(array interface{}, values ...interface{}) (interface{}, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if values == nil || arrValue.Len() == 0 {
		return array, nil
	}

	mapValue := make(map[interface{}]bool)
	for _, v := range values {
		mapValue[v] = true
	}

	result := reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := 0; i < arrValue.Len(); i++ {
		element := arrValue.Index(i)
		if !mapValue[element.Interface()] {
			result = reflect.Append(result, element)
		}
	}

	return result.Interface(), nil
}
