package array

import (
	"reflect"
)

// head returns the first element of an array (slice).
// The function accepts an array (slice) and returns the first element of the array.
// If the array is empty or not a slice, the function returns nil.
func head(array interface{}) interface{} {
	arrValue := reflect.ValueOf(array)

	if (arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array) || arrValue.Len() == 0 {
		return nil
	}

	return arrValue.Index(0).Interface()
}
