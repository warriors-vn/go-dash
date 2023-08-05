package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// lastIndexOf searches for the last occurrence of the given inputSearch value
// within the array (slice) and returns its index. The search starts from the end
// of the array or from the specified fromIndex if provided.
func lastIndexOf(array, inputSearch interface{}, fromIndex ...int) (int, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return -1, constants.ErrNotSlice
	}

	from := arrValue.Len() - 1
	if fromIndex != nil && fromIndex[0] >= 0 {
		from = fromIndex[0]
	}

	if from >= arrValue.Len() {
		from = arrValue.Len() - 1
	}

	for i := from; i >= 0; i-- {
		if reflect.DeepEqual(arrValue.Index(i).Interface(), inputSearch) {
			return i, nil
		}
	}

	return -1, nil
}
