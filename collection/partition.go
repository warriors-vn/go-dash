package collection

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// partition separates the elements of the input array into two slices based on the predicate function.
// It takes an array-like data structure and a predicate function that determines the partition.
// The predicate function should have the signature func(elementType) bool.
// The function returns two new slices: one containing elements that satisfy the predicate,
// and the other containing elements that do not satisfy the predicate.
// An error is returned if any occurs.
func partition(array, predicate interface{}) (interface{}, error) {
	arrValue, predicateValue := reflect.ValueOf(array), reflect.ValueOf(predicate)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if predicateValue.Kind() != reflect.Func {
		return nil, constants.ErrNotFunction
	}

	numParams := predicateValue.Type().NumIn()
	if numParams != 1 {
		return nil, constants.ErrNotSupport
	}

	kind, result := predicateValue.Type().In(0).Kind(), reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(array)), 0, 0)
	partitionOne, partitionTwo := reflect.MakeSlice(arrValue.Type(), 0, 0), reflect.MakeSlice(arrValue.Type(), 0, 0)
	if kind == reflect.Interface {
		return nil, constants.ErrNotSupport
	}

	for i := 0; i < arrValue.Len(); i++ {
		element := arrValue.Index(i)
		if element.Kind() != kind {
			return nil, constants.ErrIncompatible
		}

		res := predicateValue.Call([]reflect.Value{reflect.ValueOf(element.Interface())})
		if len(res) > 0 && res[0].Kind() == reflect.Bool {
			if res[0].Interface().(bool) {
				partitionOne = reflect.Append(partitionOne, element)
			} else {
				partitionTwo = reflect.Append(partitionTwo, element)
			}

		}
	}

	if partitionOne.Len() > 0 || partitionTwo.Len() > 0 {
		result = reflect.Append(result, partitionOne, partitionTwo)
	}
	return result.Interface(), nil
}
