package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// TakeRight returns a new Slice containing the Last 'number' elements from the input array.
// It takes an array-like data structure and an optional number of elements to Take.
// If 'number' is not provided or is greater than the length of the array, all elements are taken.
// The function returns the new Slice and an error if any occurs.
func TakeRight(array interface{}, number ...int) (interface{}, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if arrValue.Len() == 0 {
		return array, nil
	}

	n := 1
	if number != nil && number[0] >= 0 {
		n = number[0]
	}

	if n >= arrValue.Len() {
		n = arrValue.Len()
	}

	result := reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := arrValue.Len() - n; i < arrValue.Len(); i++ {
		result = reflect.Append(result, arrValue.Index(i))
	}

	return result.Interface(), nil
}
