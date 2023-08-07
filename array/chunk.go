package array

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Chunk divides a given array into smaller chunks of the specified size.
// The function accepts an array of any type (interface{}) and a size for each Chunk.
// It returns a Slice of slices, where each inner Slice contains elements from the original array.
// The Last Chunk may contain fewer elements if the length of the array is not divisible by the Chunk size.
func Chunk(array interface{}, size int) (interface{}, error) {
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
