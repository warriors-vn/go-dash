package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// pullAt removes elements from the input array at specified indexes and returns the modified array.
// It takes an array-like data structure and a list of indexes to remove.
// The function returns the modified array and an error if any occurs.
func pullAt(array interface{}, indexes []int) (interface{}, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	mapIndexesValue := make(map[int]int)
	for _, idx := range indexes {
		mapIndexesValue[idx]++
	}

	result := reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := 0; i < arrValue.Len(); i++ {
		element := arrValue.Index(i)
		if data, ok := mapIndexesValue[i]; ok && data > 0 {
			mapIndexesValue[i] = 0
			for j := 0; j < data; j++ {
				result = reflect.Append(result, element)
			}
		}
	}

	return result.Interface(), nil
}
