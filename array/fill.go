package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// fill elements of an array (slice) with a specified value within the specified range.
// The function accepts an array (slice), a value to fill with, and optional start and end indices.
// It returns the modified array with filled values.
// If the input is not a slice or if the start or end indices are out of bounds, the function returns an error.
func fill(array, input interface{}, start, end int) (interface{}, error) {
	arrValue, inputValue := reflect.ValueOf(array), reflect.ValueOf(input)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if start >= end || arrValue.Len() == 0 {
		return array, nil
	}

	if start < 0 {
		start = 0
	}

	if end < 0 {
		end = 1
	}

	if end > arrValue.Len() {
		end = arrValue.Len()
	}

	kind, result := arrValue.Index(0).Kind(), reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := 0; i < arrValue.Len(); i++ {
		if kind != inputValue.Kind() && kind != reflect.Interface {
			return nil, constants.ErrIncompatible
		}

		if start <= i && i < end {
			result = reflect.Append(result, inputValue)
		} else {
			result = reflect.Append(result, arrValue.Index(i))
		}
	}

	return result.Interface(), nil
}
