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

	switch kind {
	case reflect.Int:
		result, chunkPart := make([][]int, 0), make([]int, 0)
		for i := 0; i < arrValue.Len(); i++ {
			chunkPart = append(chunkPart, arrValue.Index(i).Interface().(int))
			if len(chunkPart) == size {
				result = append(result, chunkPart)
				chunkPart = make([]int, 0)
			}
		}

		if len(chunkPart) > 0 {
			result = append(result, chunkPart)
		}

		return result, nil
	case reflect.Int32:
		result, chunkPart := make([][]int32, 0), make([]int32, 0)
		for i := 0; i < arrValue.Len(); i++ {
			chunkPart = append(chunkPart, arrValue.Index(i).Interface().(int32))
			if len(chunkPart) == size {
				result = append(result, chunkPart)
				chunkPart = make([]int32, 0)
			}
		}

		if len(chunkPart) > 0 {
			result = append(result, chunkPart)
		}

		return result, nil
	case reflect.Int64:
		result, chunkPart := make([][]int64, 0), make([]int64, 0)
		for i := 0; i < arrValue.Len(); i++ {
			chunkPart = append(chunkPart, arrValue.Index(i).Interface().(int64))
			if len(chunkPart) == size {
				result = append(result, chunkPart)
				chunkPart = make([]int64, 0)
			}
		}

		if len(chunkPart) > 0 {
			result = append(result, chunkPart)
		}

		return result, nil
	case reflect.Float32:
		result, chunkPart := make([][]float32, 0), make([]float32, 0)
		for i := 0; i < arrValue.Len(); i++ {
			chunkPart = append(chunkPart, arrValue.Index(i).Interface().(float32))
			if len(chunkPart) == size {
				result = append(result, chunkPart)
				chunkPart = make([]float32, 0)
			}
		}

		if len(chunkPart) > 0 {
			result = append(result, chunkPart)
		}

		return result, nil
	case reflect.Float64:
		result, chunkPart := make([][]float64, 0), make([]float64, 0)
		for i := 0; i < arrValue.Len(); i++ {
			chunkPart = append(chunkPart, arrValue.Index(i).Interface().(float64))
			if len(chunkPart) == size {
				result = append(result, chunkPart)
				chunkPart = make([]float64, 0)
			}
		}

		if len(chunkPart) > 0 {
			result = append(result, chunkPart)
		}

		return result, nil
	case reflect.String:
		result, chunkPart := make([][]string, 0), make([]string, 0)
		for i := 0; i < arrValue.Len(); i++ {
			chunkPart = append(chunkPart, arrValue.Index(i).Interface().(string))
			if len(chunkPart) == size {
				result = append(result, chunkPart)
				chunkPart = make([]string, 0)
			}
		}

		if len(chunkPart) > 0 {
			result = append(result, chunkPart)
		}

		return result, nil
	}

	return nil, constants.ErrNotSupport
}
