package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Intersection returns a new array (Slice) containing unique elements that are present in both input arrays.
// The function accepts two arrays (slices) and returns a new array with the Intersection of unique elements.
// If either of the inputs is not a Slice, the function returns an error.
func Intersection(array, other interface{}) (interface{}, error) {
	arrValue, otherValue := reflect.ValueOf(array), reflect.ValueOf(other)

	if (arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array) || (otherValue.Kind() != reflect.Slice && otherValue.Kind() != reflect.Array) {
		return nil, constants.ErrNotSlice
	}

	if arrValue.Len() == 0 {
		return array, nil
	}

	if otherValue.Len() == 0 {
		return other, nil
	}

	mapOtherValue := make(map[interface{}]int)
	for i := 0; i < otherValue.Len(); i++ {
		mapOtherValue[otherValue.Index(i).Interface()]++
	}

	result := reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := 0; i < arrValue.Len(); i++ {
		element := arrValue.Index(i)
		if data, ok := mapOtherValue[element.Interface()]; ok && data != 0 {
			mapOtherValue[element.Interface()] = 0
			result = reflect.Append(result, element)
		}
	}

	return result.Interface(), nil
}
