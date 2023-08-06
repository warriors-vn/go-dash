package collection

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// goDashMap applies the iteratee function to each element in the input array and returns a new slice of results.
// It takes an array-like data structure and an iteratee function that operates on each element.
// The iteratee function should have the signature func(elementType) mappedElementType.
// The function returns a new slice of mapped elements and an error if any occurs.
func goDashMap(array, iteratee interface{}) (interface{}, error) {
	arrValue, iterateeValue := reflect.ValueOf(array), reflect.ValueOf(iteratee)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if iterateeValue.Kind() != reflect.Func {
		return nil, constants.ErrNotFunction
	}

	numParams := iterateeValue.Type().NumIn()
	if numParams != 1 {
		return nil, constants.ErrNotSupport
	}

	kind, result := iterateeValue.Type().In(0).Kind(), reflect.MakeSlice(arrValue.Type(), 0, 0)
	if kind == reflect.Interface {
		return nil, constants.ErrNotSupport
	}
	for i := 0; i < arrValue.Len(); i++ {
		res := iterateeValue.Call([]reflect.Value{reflect.ValueOf(arrValue.Index(i).Interface())})
		if len(res) > 0 && res[0].Kind() == kind {
			result = reflect.Append(result, res[0])
		}
	}

	return result.Interface(), nil
}
