package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Tail returns a new Slice containing all elements of the input array except the first one.
// It takes an array-like data structure as input and returns a new Slice excluding the first element.
// The function returns the new Slice and an error if any occurs.
func Tail(array interface{}) (interface{}, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if arrValue.Len() == 0 {
		return array, nil
	}

	result := reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := 1; i < arrValue.Len(); i++ {
		result = reflect.Append(result, arrValue.Index(i))
	}

	return result.Interface(), nil
}
