package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Reverse reverses the elements in the input array.
// It takes an array-like data structure and returns the modified array with reversed elements.
// The function returns the modified array and an error if any occurs.
func Reverse(array interface{}) (interface{}, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	result := reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := arrValue.Len() - 1; i >= 0; i-- {
		result = reflect.Append(result, arrValue.Index(i))
	}

	return result.Interface(), nil
}
