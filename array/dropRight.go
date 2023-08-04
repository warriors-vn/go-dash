package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// dropRight returns a new array with `num` elements dropped from the end.
// The function accepts an array of any type (interface{}) and an optional number `num` of elements to drop.
// It returns a new array with the specified number of elements removed from the end.
func dropRight(array interface{}, num ...int) (interface{}, error) {
	arrValue, n := reflect.ValueOf(array), 0

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

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if arrValue.Len() == 0 {
		return array, nil
	}

	kind := arrValue.Index(0).Kind()
	if kind == reflect.Interface {
		return nil, constants.ErrNotSupport
	}

	switch kind {
	case reflect.Int:
		result := make([]int, 0)
		if n >= arrValue.Len() {
			return result, nil
		}

		for i := 0; i < arrValue.Len()-n; i++ {
			element := arrValue.Index(i)
			result = append(result, element.Interface().(int))
		}

		return result, nil
	case reflect.Int32:
		result := make([]int32, 0)
		if n >= arrValue.Len() {
			return result, nil
		}

		for i := 0; i < arrValue.Len()-n; i++ {
			element := arrValue.Index(i)
			result = append(result, element.Interface().(int32))
		}

		return result, nil
	case reflect.Int64:
		result := make([]int64, 0)
		if n >= arrValue.Len() {
			return result, nil
		}

		for i := 0; i < arrValue.Len()-n; i++ {
			element := arrValue.Index(i)
			result = append(result, element.Interface().(int64))
		}

		return result, nil
	case reflect.Float32:
		result := make([]float32, 0)
		if n >= arrValue.Len() {
			return result, nil
		}

		for i := 0; i < arrValue.Len()-n; i++ {
			element := arrValue.Index(i)
			result = append(result, element.Interface().(float32))
		}

		return result, nil
	case reflect.Float64:
		result := make([]float64, 0)
		if n >= arrValue.Len() {
			return result, nil
		}

		for i := 0; i < arrValue.Len()-n; i++ {
			element := arrValue.Index(i)
			result = append(result, element.Interface().(float64))
		}

		return result, nil
	case reflect.String:
		result := make([]string, 0)
		if n >= arrValue.Len() {
			return result, nil
		}

		for i := 0; i < arrValue.Len()-n; i++ {
			element := arrValue.Index(i)
			result = append(result, element.Interface().(string))
		}

		return result, nil
	}

	return nil, constants.ErrNotSupport
}
