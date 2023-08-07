package go_math

import "reflect"

// Subtract performs subtraction on two values of various numeric types.
// It returns the result as an interface{} to handle different numeric types.
func Subtract(minuend, subtrahend interface{}) interface{} {
	typeOfMinuend, typeOfSubtrahend := reflect.TypeOf(minuend), reflect.TypeOf(subtrahend)

	kind := typeOfMinuend.Kind()
	if typeOfMinuend.Kind() != typeOfSubtrahend.Kind() {
		return 0
	}

	switch kind {
	case reflect.Int:
		if result, ok := subtrahend.(int); ok {
			return minuend.(int) - result
		}
	case reflect.Int32:
		if result, ok := subtrahend.(int32); ok {
			return minuend.(int32) - result
		}
	case reflect.Int64:
		if result, ok := subtrahend.(int64); ok {
			return minuend.(int64) - result
		}
	case reflect.Float32:
		if result, ok := subtrahend.(float32); ok {
			return minuend.(float32) - result
		}
	case reflect.Float64:
		if result, ok := subtrahend.(float64); ok {
			return minuend.(float64) - result
		}
	}

	return 0
}
