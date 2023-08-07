package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Drop returns a new array with `n` elements dropped from the beginning.
// The function accepts an array of any type (interface{}) and an optional number `num` of elements to Drop.
// It returns a new array with the specified number of elements removed from the beginning.
func Drop(array interface{}, num ...int) (interface{}, error) {
	arrValue, n := reflect.ValueOf(array), 0

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if num != nil {
		n = num[0]
	} else {
		return array, nil
	}

	if n == 0 || arrValue.Len() == 0 {
		return array, nil
	}

	if n < 0 {
		return nil, constants.ErrParamLessThanZero
	}

	result := reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := n; i < arrValue.Len(); i++ {
		result = reflect.Append(result, arrValue.Index(i))
	}

	return result.Interface(), nil
}
