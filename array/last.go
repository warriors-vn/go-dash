package array

import "reflect"

// last returns the last element of an array (slice).
// The function accepts an array (slice) and returns its last element.
// If the input is not a slice or if the array is empty, the function returns nil.
func last(array interface{}) interface{} {
	arrValue := reflect.ValueOf(array)

	if (arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array) || arrValue.Len() == 0 {
		return nil
	}

	return arrValue.Index(arrValue.Len() - 1).Interface()
}
