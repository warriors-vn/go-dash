package go_math

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// MaxInt finds and returns the maximum value from the given slice of integers 'i'.
func MaxInt(i []int) int {
	if len(i) == 0 {
		return 0
	}

	max := i[0]
	for _, value := range i {
		if value > max {
			max = value
		}
	}

	return max
}

// MaxInt32 finds and returns the maximum value from the given slice of int32 values 'i'.
func MaxInt32(i []int32) int32 {
	if len(i) == 0 {
		return 0
	}

	max := i[0]

	for _, value := range i {
		if value > max {
			max = value
		}
	}

	return max
}

// MaxInt64 finds and returns the maximum value from the given slice of int64 values 'i'.
func MaxInt64(i []int64) int64 {
	if len(i) == 0 {
		return 0
	}

	max := i[0]

	for _, value := range i {
		if value > max {
			max = value
		}
	}

	return max
}

// MaxFloat32 finds and returns the maximum value from the given slice of float32 values 'i'.
func MaxFloat32(i []float32) float32 {
	if len(i) == 0 {
		return 0
	}

	max := i[0]

	for _, value := range i {
		if value > max {
			max = value
		}
	}

	return max
}

// MaxFloat64 finds and returns the maximum value from the given slice of float64 values 'i'.
func MaxFloat64(i []float64) float64 {
	if len(i) == 0 {
		return 0
	}

	max := i[0]

	for _, value := range i {
		if value > max {
			max = value
		}
	}

	return max
}

// MaxField finds and returns the maximum value of a specified field in a list of structs.
// The 'list' parameter is the slice of structs, and 'fieldName' specifies the field to compare.
// The function returns the maximum field value as an interface{} and an error if any issue occurs.
func MaxField(list interface{}, fieldName string) (interface{}, error) {
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

	maxValue := fieldFirst.Interface()

	for i := 1; i < value.Len(); i++ {
		elem := value.Index(i)
		field := elem.FieldByName(fieldName)

		if field.Interface() == nil {
			continue
		}

		if reflect.ValueOf(maxValue).Type() != field.Type() {
			return nil, constants.ErrIncompatible
		}

		if field.Interface().(int) > maxValue.(int) {
			maxValue = field.Interface()
		}
	}

	return maxValue, nil
}
