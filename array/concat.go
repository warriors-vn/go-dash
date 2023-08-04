package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// concat concatenates two arrays into a single array.
// The function accepts two arrays of any type (interface{}) and returns a single concatenated array.
func concat(array, extend interface{}) (interface{}, error) {
	arrValue, extendValue := reflect.ValueOf(array), reflect.ValueOf(extend)

	if (arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array) ||
		(extendValue.Kind() != reflect.Slice && extendValue.Kind() != reflect.Array) {
		return nil, constants.ErrNotSlice
	}

	if arrValue.Len() == 0 {
		return extend, nil
	}

	if extendValue.Len() == 0 {
		return array, nil
	}

	kind := arrValue.Index(0).Kind()
	if kind == reflect.Interface {
		return nil, constants.ErrNotSupport
	}

	switch kind {
	case reflect.Int:
		result := make([]int, 0)
		for i := 0; i < arrValue.Len(); i++ {
			element := arrValue.Index(i)
			result = append(result, element.Interface().(int))
		}

		for i := 0; i < extendValue.Len(); i++ {
			element := extendValue.Index(i)
			if kind != element.Kind() {
				return nil, constants.ErrIncompatible
			}
			result = append(result, element.Interface().(int))
		}

		return result, nil
	case reflect.Int32:
		result := make([]int32, 0)
		for i := 0; i < arrValue.Len(); i++ {
			element := arrValue.Index(i)
			result = append(result, element.Interface().(int32))
		}

		for i := 0; i < extendValue.Len(); i++ {
			element := extendValue.Index(i)
			if kind != element.Kind() {
				return nil, constants.ErrIncompatible
			}
			result = append(result, element.Interface().(int32))
		}

		return result, nil
	case reflect.Int64:
		result := make([]int64, 0)
		for i := 0; i < arrValue.Len(); i++ {
			element := arrValue.Index(i)
			result = append(result, element.Interface().(int64))
		}

		for i := 0; i < extendValue.Len(); i++ {
			element := extendValue.Index(i)
			if kind != element.Kind() {
				return nil, constants.ErrIncompatible
			}
			result = append(result, element.Interface().(int64))
		}

		return result, nil
	case reflect.Float32:
		result := make([]float32, 0)
		for i := 0; i < arrValue.Len(); i++ {
			element := arrValue.Index(i)
			result = append(result, element.Interface().(float32))
		}

		for i := 0; i < extendValue.Len(); i++ {
			element := extendValue.Index(i)
			if kind != element.Kind() {
				return nil, constants.ErrIncompatible
			}
			result = append(result, element.Interface().(float32))
		}

		return result, nil
	case reflect.Float64:
		result := make([]float64, 0)
		for i := 0; i < arrValue.Len(); i++ {
			element := arrValue.Index(i)
			result = append(result, element.Interface().(float64))
		}

		for i := 0; i < extendValue.Len(); i++ {
			element := extendValue.Index(i)
			if kind != element.Kind() {
				return nil, constants.ErrIncompatible
			}
			result = append(result, element.Interface().(float64))
		}

		return result, nil
	case reflect.String:
		result := make([]string, 0)
		for i := 0; i < arrValue.Len(); i++ {
			element := arrValue.Index(i)
			result = append(result, element.Interface().(string))
		}

		for i := 0; i < extendValue.Len(); i++ {
			element := extendValue.Index(i)
			if kind != element.Kind() {
				return nil, constants.ErrIncompatible
			}
			result = append(result, element.Interface().(string))
		}

		return result, nil
	}

	return nil, constants.ErrNotSupport
}
