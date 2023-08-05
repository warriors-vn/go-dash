package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// nth returns the element at the specified index from the array (slice).
// The index is 1-based, so the first element is at index 1.
func nth(array interface{}, number int) (interface{}, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if number >= arrValue.Len() {
		return nil, nil
	}

	if number < 0 {
		if number+arrValue.Len() < 0 {
			return nil, nil
		}
		number = number + arrValue.Len()
	}

	return arrValue.Index(number).Interface(), nil
}
