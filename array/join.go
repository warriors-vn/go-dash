package array

import (
	"fmt"
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// join takes an array (slice) of elements and a separator string,
// and returns a single string by joining the elements with the separator.
func join(array interface{}, separator string) (string, error) {
	arrValue, result := reflect.ValueOf(array), ""

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return result, constants.ErrNotSlice
	}

	for i := 0; i < arrValue.Len(); i++ {
		if i == arrValue.Len()-1 {
			result += fmt.Sprintf("%v", arrValue.Index(i).Interface())
		} else {
			result += fmt.Sprintf("%v%v", arrValue.Index(i).Interface(), separator)
		}
	}

	return result, nil
}
