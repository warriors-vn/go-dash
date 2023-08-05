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

	mapExcludeValue := make(map[interface{}]int)
	for i := 0; i < excludeValue.Len(); i++ {
		mapExcludeValue[excludeValue.Index(i).Interface()]++
	}

	result := reflect.MakeSlice(arrValue.Type(), 0, 0)
	for i := 0; i < arrValue.Len(); i++ {
		element := arrValue.Index(i)
		if _, ok := mapExcludeValue[element.Interface()]; !ok {
			result = reflect.Append(result, element)
		}
	}

	return result.Interface(), nil
}
