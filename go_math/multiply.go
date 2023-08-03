package go_math

import (
	"reflect"

	"go-dash/constants"
)

// multiply calculates the product of the multiplier and subtrahend.
// The function accepts two arguments, multiplier and subtrahend, both of type interface{}.
// It returns the result of the multiplication as an interface{} and an error if any issue occurs.
func multiply(multiplier, subtrahend interface{}) (interface{}, error) {
	typeOfMultiplier, typeOfSubtrahend := reflect.TypeOf(multiplier), reflect.TypeOf(subtrahend)

	kind := typeOfMultiplier.Kind()
	if typeOfMultiplier.Kind() != typeOfSubtrahend.Kind() {
		return 0, constants.ErrIncompatible
	}

	switch kind {
	case reflect.Int:
		if result, ok := subtrahend.(int); ok {
			return multiplier.(int) * result, nil
		}
	case reflect.Int32:
		if result, ok := subtrahend.(int32); ok {
			return multiplier.(int32) * result, nil
		}
	case reflect.Int64:
		if result, ok := subtrahend.(int64); ok {
			return multiplier.(int64) * result, nil
		}
	case reflect.Float32:
		if result, ok := subtrahend.(float32); ok {
			return multiplier.(float32) * result, nil
		}
	case reflect.Float64:
		if result, ok := subtrahend.(float64); ok {
			return multiplier.(float64) * result, nil
		}
	}

	return 0, constants.ErrNotSupport
}
