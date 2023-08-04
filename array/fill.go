package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// fill elements of an array (slice) with a specified value within the specified range.
// The function accepts an array (slice), a value to fill with, and optional start and end indices.
// It returns the modified array with filled values.
// If the input is not a slice or if the start or end indices are out of bounds, the function returns an error.
func fill(array, input interface{}, start, end int) (interface{}, error) {
	arrValue, inputValue := reflect.ValueOf(array), reflect.ValueOf(input)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if start >= end || arrValue.Len() == 0 {
		return array, nil
	}

	if start < 0 {
		start = 0
	}

	if end < 0 {
		end = 1
	}

	if end > arrValue.Len() {
		end = arrValue.Len()
	}

	kind := arrValue.Index(0).Kind()
	if kind == reflect.Interface {
		return nil, constants.ErrNotSupport
	}

	switch kind {
	case reflect.Int:
		result := make([]int, 0)
		for i := 0; i < arrValue.Len(); i++ {
			if kind != inputValue.Kind() {
				return nil, constants.ErrIncompatible
			}

			if start <= i && i < end {
				result = append(result, inputValue.Interface().(int))
			} else {
				result = append(result, arrValue.Index(i).Interface().(int))
			}
		}

		return result, nil
	case reflect.Int32:
		result := make([]int32, 0)
		for i := 0; i < arrValue.Len(); i++ {
			if kind != inputValue.Kind() {
				return nil, constants.ErrIncompatible
			}

			if start <= i && i < end {
				result = append(result, inputValue.Interface().(int32))
			} else {
				result = append(result, arrValue.Index(i).Interface().(int32))
			}
		}

		return result, nil
	case reflect.Int64:
		result := make([]int64, 0)
		for i := 0; i < arrValue.Len(); i++ {
			if kind != inputValue.Kind() {
				return nil, constants.ErrIncompatible
			}

			if start <= i && i < end {
				result = append(result, inputValue.Interface().(int64))
			} else {
				result = append(result, arrValue.Index(i).Interface().(int64))
			}
		}

		return result, nil
	case reflect.Float32:
		result := make([]float32, 0)
		for i := 0; i < arrValue.Len(); i++ {
			if kind != inputValue.Kind() {
				return nil, constants.ErrIncompatible
			}

			if start <= i && i < end {
				result = append(result, inputValue.Interface().(float32))
			} else {
				result = append(result, arrValue.Index(i).Interface().(float32))
			}
		}

		return result, nil
	case reflect.Float64:
		result := make([]float64, 0)
		for i := 0; i < arrValue.Len(); i++ {
			if kind != inputValue.Kind() {
				return nil, constants.ErrIncompatible
			}

			if start <= i && i < end {
				result = append(result, inputValue.Interface().(float64))
			} else {
				result = append(result, arrValue.Index(i).Interface().(float64))
			}
		}

		return result, nil
	case reflect.String:
		result := make([]string, 0)
		for i := 0; i < arrValue.Len(); i++ {
			if kind != inputValue.Kind() {
				return nil, constants.ErrIncompatible
			}

			if start <= i && i < end {
				result = append(result, inputValue.Interface().(string))
			} else {
				result = append(result, arrValue.Index(i).Interface().(string))
			}
		}

		return result, nil
	}

	return nil, constants.ErrNotSupport
}
