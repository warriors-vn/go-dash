package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// sortedUniq returns a new slice with only unique elements from the sorted input slice.
// It takes an array-like data structure as input, assumes that the input slice is sorted,
// and returns a new slice containing only unique elements.
// The function returns the new slice and an error if any occurs.
func sortedUniq(array interface{}) (interface{}, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	mapArrValue := make(map[interface{}]int)
	for i := 0; i < arrValue.Len(); i++ {
		mapArrValue[arrValue.Index(i).Interface()]++
	}

	result := reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := 0; i < arrValue.Len(); i++ {
		element := arrValue.Index(i)
		if data, ok := mapArrValue[element.Interface()]; ok && data > 0 {
			mapArrValue[element.Interface()] = 0
			result = reflect.Append(result, element)
		}
	}

	return result.Interface(), nil
}
