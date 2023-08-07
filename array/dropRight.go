package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// DropRight returns a new array with `num` elements dropped from the end.
// The function accepts an array of any type (interface{}) and an optional number `num` of elements to Drop.
// It returns a new array with the specified number of elements removed from the end.
func DropRight(array interface{}, num ...int) (interface{}, error) {
	arrValue, n := reflect.ValueOf(array), 0

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if num != nil {
		n = num[0]
	} else {
		return array, nil
	}

	if n == 0 {
		return array, nil
	}

	if n < 0 {
		return nil, constants.ErrParamLessThanZero
	}

	if arrValue.Len() == 0 {
		return array, nil
	}

	result := reflect.MakeSlice(arrValue.Type(), 0, 0)
	if n >= arrValue.Len() {
		return result.Interface(), nil
	}

	for i := 0; i < arrValue.Len()-n; i++ {
		result = reflect.Append(result, arrValue.Index(i))
	}

	return result.Interface(), nil
}
