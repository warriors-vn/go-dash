package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// PullAll removes all occurrences of specified elements from the array (Slice).
func PullAll(array, removes interface{}) (interface{}, error) {
	arrValue, removesValue := reflect.ValueOf(array), reflect.ValueOf(removes)

	if (arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array) ||
		(removesValue.Kind() != reflect.Slice && removesValue.Kind() != reflect.Array) {
		return nil, constants.ErrNotSlice
	}

	mapArrValue := make(map[interface{}]int)
	for i := 0; i < arrValue.Len(); i++ {
		mapArrValue[arrValue.Index(i).Interface()]++
	}

	for i := 0; i < removesValue.Len(); i++ {
		element := removesValue.Index(i)
		if _, ok := mapArrValue[element.Interface()]; ok {
			mapArrValue[element.Interface()] = 0
		}
	}

	result := reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := 0; i < arrValue.Len(); i++ {
		element := arrValue.Index(i)
		if data, ok := mapArrValue[element.Interface()]; ok && data > 0 {
			result = reflect.Append(result, element)
		}
	}

	return result.Interface(), nil
}
