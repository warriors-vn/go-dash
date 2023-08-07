package lang

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// Lte checks if the first value is less than or equal to the second value.
// The function accepts two values of any comparable type (int, float, string, etc.) and returns true if the first value is less than or equal to the second value, false otherwise.
// If the values are not comparable, the function returns an error.
func Lte(value, other interface{}) (bool, error) {
	valueOf, otherOfValue := reflect.ValueOf(value), reflect.ValueOf(other)

	if valueOf.Kind() != otherOfValue.Kind() {
		return false, constants.ErrIncompatible
	}

	switch valueOf.Kind() {
	case reflect.Int:
		return valueOf.Interface().(int) <= otherOfValue.Interface().(int), nil
	case reflect.Int32:
		return valueOf.Interface().(int32) <= otherOfValue.Interface().(int32), nil
	case reflect.Int64:
		return valueOf.Interface().(int64) <= otherOfValue.Interface().(int64), nil
	case reflect.Float32:
		return valueOf.Interface().(float32) <= otherOfValue.Interface().(float32), nil
	case reflect.Float64:
		return valueOf.Interface().(float64) <= otherOfValue.Interface().(float64), nil
	}

	return false, constants.ErrNotSupport
}
