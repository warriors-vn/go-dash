package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Initial returns a new array (Slice) containing all elements of the input array except the Last one.
// The function accepts an array (Slice) and returns a new array with all elements except the Last one.
// If the input is not a Slice or if the array is empty, the function returns an error.
func Initial(array interface{}) (interface{}, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if arrValue.Len() == 0 {
		return array, nil
	}

	result := reflect.MakeSlice(arrValue.Type(), arrValue.Len()-1, arrValue.Len()-1)

	for i := 0; i < arrValue.Len()-1; i++ {
		result.Index(i).Set(arrValue.Index(i))
	}

	return result.Interface(), nil
}
