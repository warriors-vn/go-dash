package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Slice extracts a sub-Slice from the input array, starting from the 'start' index (inclusive)
// up to the 'end' index (exclusive).
// It takes an array-like data structure, the start index, and the end index as arguments.
// The function returns the sub-Slice and an error if any occurs.
func Slice(array interface{}, start, end int) (interface{}, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if start < 0 || end < 0 {
		return nil, constants.ErrParamLessThanZero
	}

	if start >= end || start >= arrValue.Len() || end > arrValue.Len() {
		return nil, constants.ErrOutOfRange
	}

	result := reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := start; i < end; i++ {
		result = reflect.Append(result, arrValue.Index(i))
	}

	return result.Interface(), nil
}
