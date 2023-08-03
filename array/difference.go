package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

func difference(array interface{}, exclude ...interface{}) (interface{}, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	var excludeArr interface{}
	if exclude == nil {
		return array, nil
	} else {
		excludeArr = exclude[0]
	}

	excludeValue := reflect.ValueOf(excludeArr)
	if excludeValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if excludeValue.Len() == 0 || arrValue.Len() == 0 {
		return array, nil
	}

	kind, mapArrayValue := arrValue.Index(0).Kind(), make(map[interface{}]int)
	if kind == reflect.Interface {
		return nil, constants.ErrNotSupport
	}

	for i := 0; i < arrValue.Len(); i++ {
		element := arrValue.Index(i)
		mapArrayValue[element.Interface()]++
	}

	for i := 0; i < excludeValue.Len(); i++ {
		element := excludeValue.Index(i)
		if element.Kind() != kind {
			return nil, constants.ErrIncompatible
		}

		if _, ok := mapArrayValue[element.Interface()]; ok {
			mapArrayValue[element.Interface()] = 0
		}
	}

	switch kind {
	case reflect.Int:
		result := make([]int, 0)
		for i := 0; i < arrValue.Len(); i++ {
			element := arrValue.Index(i)
			if data, ok := mapArrayValue[element.Interface()]; ok && data != 0 {
				result = append(result, element.Interface().(int))
			}
		}

		return result, nil
	case reflect.Int32:
		result := make([]int32, 0)
		for i := 0; i < arrValue.Len(); i++ {
			element := arrValue.Index(i)
			if data, ok := mapArrayValue[element.Interface()]; ok && data != 0 {
				result = append(result, element.Interface().(int32))
			}
		}

		return result, nil
	case reflect.Int64:
		result := make([]int64, 0)
		for i := 0; i < arrValue.Len(); i++ {
			element := arrValue.Index(i)
			if data, ok := mapArrayValue[element.Interface()]; ok && data != 0 {
				result = append(result, element.Interface().(int64))
			}
		}

		return result, nil
	case reflect.Float32:
		result := make([]float32, 0)
		for i := 0; i < arrValue.Len(); i++ {
			element := arrValue.Index(i)
			if data, ok := mapArrayValue[element.Interface()]; ok && data != 0 {
				result = append(result, element.Interface().(float32))
			}
		}

		return result, nil
	case reflect.Float64:
		result := make([]float64, 0)
		for i := 0; i < arrValue.Len(); i++ {
			element := arrValue.Index(i)
			if data, ok := mapArrayValue[element.Interface()]; ok && data != 0 {
				result = append(result, element.Interface().(float64))
			}
		}

		return result, nil
	case reflect.Bool:
		result := make([]bool, 0)
		for i := 0; i < arrValue.Len(); i++ {
			element := arrValue.Index(i)
			if data, ok := mapArrayValue[element.Interface()]; ok && data != 0 {
				result = append(result, element.Interface().(bool))
			}
		}

		return result, nil
	case reflect.String:
		result := make([]string, 0)
		for i := 0; i < arrValue.Len(); i++ {
			element := arrValue.Index(i)
			if data, ok := mapArrayValue[element.Interface()]; ok && data != 0 {
				result = append(result, element.Interface().(string))
			}
		}

		return result, nil
	}

	return nil, constants.ErrNotSupport
}
