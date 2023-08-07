package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// SortedUniq returns a new Slice with only unique elements from the sorted input Slice.
// It takes an array-like data structure as input, assumes that the input Slice is sorted,
// and returns a new Slice containing only unique elements.
// The function returns the new Slice and an error if any occurs.
func SortedUniq(array interface{}) (interface{}, error) {
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
