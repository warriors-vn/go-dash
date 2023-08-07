package collection

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Includes checks if the input array contains the specified search value.
// It takes an array-like data structure, a search value, and an optional starting index.
// If the starting index is provided, the search starts from that index.
// The function returns true if the search is found, false otherwise, or an error if any occurs.
func Includes(array interface{}, search interface{}, fromIndex ...int) (interface{}, error) {
	arrValue, searchValue := reflect.ValueOf(array), reflect.ValueOf(search)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return false, constants.ErrNotSlice
	}

	from := 0
	if fromIndex != nil && fromIndex[0] > 0 {
		from = fromIndex[0]
	}

	if from >= arrValue.Len() {
		return false, nil
	}

	for i := from; i < arrValue.Len(); i++ {
		if reflect.DeepEqual(arrValue.Index(i).Interface(), searchValue.Interface()) {
			return true, nil
		}
	}

	return false, nil
}
