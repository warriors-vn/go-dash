package lang

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// castArray converts a value to an array if it is not already an array.
// The function accepts a value of any type (interface{}) and returns an array containing the value.
// If the value is already an array, it is returned as-is.
func castArray(value interface{}) (interface{}, error) {
	switch reflect.ValueOf(value).Kind() {
	case reflect.Int:
		return []int{value.(int)}, nil
	case reflect.Int32:
		return []int32{value.(int32)}, nil
	case reflect.Int64:
		return []int64{value.(int64)}, nil
	case reflect.Float32:
		return []float32{value.(float32)}, nil
	case reflect.Float64:
		return []float64{value.(float64)}, nil
	case reflect.String:
		return []string{value.(string)}, nil
	case reflect.Bool:
		return []bool{value.(bool)}, nil
	}

	return nil, constants.ErrNotSupport
}
