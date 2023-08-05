package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// chunk divides a given array into smaller chunks of the specified size.
// The function accepts an array of any type (interface{}) and a size for each chunk.
// It returns a slice of slices, where each inner slice contains elements from the original array.
// The last chunk may contain fewer elements if the length of the array is not divisible by the chunk size.
func chunk(array interface{}, size int) (interface{}, error) {
	arrValue := reflect.ValueOf(array)

	if size < 0 {
		return nil, constants.ErrParamLessThanZero
	}

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	if arrValue.Len() == 0 || size == 0 {
		return nil, constants.ErrEmptyList
	}

	kind := arrValue.Index(0).Kind()
	if kind == reflect.Interface {
		return nil, constants.ErrNotSupport
	}

	chunkPart := reflect.MakeSlice(arrValue.Type(), 0, 0)
	result := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(array)), 0, 0)
	for i := 0; i < arrValue.Len(); i++ {
		chunkPart = reflect.Append(chunkPart, arrValue.Index(i))
		if chunkPart.Len() == size {
			result = reflect.Append(result, chunkPart)
			chunkPart = reflect.MakeSlice(arrValue.Type(), 0, 0)
		}
	}

	if chunkPart.Len() > 0 {
		result = reflect.Append(result, chunkPart)
	}

	return result.Interface(), nil
}
