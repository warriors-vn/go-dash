package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Xor returns a new Slice that contains the elements that appear in an odd number of input arrays.
// It takes variadic arguments representing array-like data structures and returns a new Slice
// containing elements that are present in an odd number of input arrays.
// The function returns the new Slice and an error if any occurs.
func Xor(arrays ...interface{}) (interface{}, error) {
	if arrays == nil {
		return nil, constants.ErrNotSupport
	}

	arrFirstValueOf := reflect.ValueOf(arrays[0])
	if arrFirstValueOf.Kind() != reflect.Slice && arrFirstValueOf.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if arrFirstValueOf.Len() == 0 {
		return nil, constants.ErrNotSupport
	}

	arrMapValue := make([]map[interface{}]bool, 0)
	for i := 0; i < len(arrays); i++ {
		arr := reflect.ValueOf(arrays[i])
		if arr.Kind() != reflect.Slice && arr.Kind() != reflect.Array {
			continue
		}

		mapValue := make(map[interface{}]bool)
		for j := 0; j < arr.Len(); j++ {
			mapValue[arr.Index(j).Interface()] = true
		}

		arrMapValue = append(arrMapValue, mapValue)
	}

	result, firstKind := reflect.MakeSlice(arrFirstValueOf.Type(), 0, 0), arrFirstValueOf.Index(0).Kind()
	for i := 0; i < len(arrays); i++ {
		arr := reflect.ValueOf(arrays[i])
		if arr.Kind() != reflect.Slice && arr.Kind() != reflect.Array {
			continue
		}

		for j := 0; j < arr.Len(); j++ {
			element, isSetTrue := arr.Index(j), false
			for _, mapValue := range arrMapValue {
				if isSetTrue {
					mapValue[element.Interface()] = true
				}

				if !mapValue[element.Interface()] && !isSetTrue {
					if firstKind != element.Kind() {
						return nil, constants.ErrIncompatible
					}
					isSetTrue = true
					result = reflect.Append(result, element)
				}
			}
		}
	}

	return result.Interface(), nil
}
