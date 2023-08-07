package array

import (
	"reflect"
)

// Head returns the first element of an array (Slice).
// The function accepts an array (Slice) and returns the first element of the array.
// If the array is empty or not a Slice, the function returns nil.
func Head(array interface{}) interface{} {
	arrValue := reflect.ValueOf(array)

	if (arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array) || arrValue.Len() == 0 {
		return nil
	}

	return arrValue.Index(0).Interface()
}
