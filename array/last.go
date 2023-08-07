package array

import "reflect"

// Last returns the Last element of an array (Slice).
// The function accepts an array (Slice) and returns its Last element.
// If the input is not a Slice or if the array is empty, the function returns nil.
func Last(array interface{}) interface{} {
	arrValue := reflect.ValueOf(array)

	if (arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array) || arrValue.Len() == 0 {
		return nil
	}

	return arrValue.Index(arrValue.Len() - 1).Interface()
}
