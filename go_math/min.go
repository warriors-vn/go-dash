package go_math

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// minInt finds and returns the minimum value from the given slice of integers 'i'.
func minInt(i []int) int {
	if len(i) == 0 {
		return 0
	}

	min := i[0]
	for _, value := range i {
		if value < min {
			min = value
		}
	}

	return min
}

// minInt32 finds and returns the minimum value from the given slice of int32 values 'i'.
func minInt32(i []int32) int32 {
	if len(i) == 0 {
		return 0
	}

	min := i[0]

	for _, value := range i {
		if value < min {
			min = value
		}
	}

	return min
}

// minInt64 finds and returns the minimum value from the given slice of int64 values 'i'.
func minInt64(i []int64) int64 {
	if len(i) == 0 {
		return 0
	}

	min := i[0]

	for _, value := range i {
		if value < min {
			min = value
		}
	}

	return min
}

// minFloat32 finds and returns the minimum value from the given slice of float32 values 'i'.
func minFloat32(i []float32) float32 {
	if len(i) == 0 {
		return 0
	}

	min := i[0]

	for _, value := range i {
		if value < min {
			min = value
		}
	}

	return min
}

// minFloat64 finds and returns the minimum value from the given slice of float64 values 'i'.
func minFloat64(i []float64) float64 {
	if len(i) == 0 {
		return 0
	}

	min := i[0]

	for _, value := range i {
		if value < min {
			min = value
		}
	}

	return min
}

// minField finds and returns the minimum value of a specified field in a list of structs.
// The 'list' parameter is the slice of structs, and 'fieldName' specifies the field to compare.
// The function returns the minimum field value as an interface{} and an error if any issue occurs.
func minField(list interface{}, fieldName string) (interface{}, error) {
	value := reflect.ValueOf(list)
	if value.Kind() != reflect.Slice {
		return nil, constants.ErrNotSlice
	}

	if value.Len() == 0 {
		return nil, constants.ErrEmptyList
	}

	fieldFirst := value.Index(0).FieldByName(fieldName)
	if !fieldFirst.IsValid() {
		return nil, constants.ErrFieldNotFound
	}

	minValue := fieldFirst.Interface()

	for i := 1; i < value.Len(); i++ {
		elem := value.Index(i)
		field := elem.FieldByName(fieldName)

		if field.Interface() == nil {
			continue
		}

		if reflect.ValueOf(minValue).Type() != field.Type() {
			return nil, constants.ErrIncompatible
		}

		if field.Interface().(int) < minValue.(int) {
			minValue = field.Interface()
		}
	}

	return minValue, nil
}
