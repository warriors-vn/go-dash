package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// IsContainsArray checks if a given element is present in the array.
// The function accepts an array (Slice) and an element of any type (interface{}) and returns true if the element is found in the array, false otherwise.
// If the array or element is of an unsupported type, the function returns an error.
func IsContainsArray(array, element interface{}) (bool, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return false, constants.ErrNotSlice
	}

	for i := 0; i < arrValue.Len(); i++ {
		if reflect.DeepEqual(element, arrValue.Index(i).Interface()) {
			return true, nil
		}
	}

	return false, nil
}
