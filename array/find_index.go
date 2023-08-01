package array

import (
	"reflect"
)

// findIndex uses reflection to find the index of the 'target' value in the 'array'.
// It returns the index of the first occurrence of 'target' and a boolean indicating if the target was found.
// If 'target' is not found in 'array', the function returns -1 and false.
func findIndex(array, target interface{}) (int, bool) {
	arrValue := reflect.ValueOf(array)
	targetValue := reflect.ValueOf(target)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return -1, false
	}

	for i := 0; i < arrValue.Len(); i++ {
		elem := arrValue.Index(i)
		if reflect.DeepEqual(elem.Interface(), targetValue.Interface()) {
			return i, true
		}
	}

	return -1, false
}
