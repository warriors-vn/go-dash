package go_math

import (
	"reflect"
)

// add performs a mathematical operation on the provided 'augend' and 'addend'.
// The function supports addition of integer and floating-point values.
// It returns the result of the mathematical operation as an interface{} value.
func add(augend, addend interface{}) interface{} {
	typeOfAugend, typeOfAddend := reflect.TypeOf(augend), reflect.TypeOf(addend)

	kind := typeOfAugend.Kind()
	if typeOfAugend.Kind() != typeOfAddend.Kind() {
		return 0
	}

	switch kind {
	case reflect.Int:
		if result, ok := addend.(int); ok {
			return augend.(int) + result
		}
	case reflect.Int32:
		if result, ok := addend.(int32); ok {
			return augend.(int32) + result
		}
	case reflect.Int64:
		if result, ok := addend.(int64); ok {
			return augend.(int64) + result
		}
	case reflect.Float32:
		if result, ok := addend.(float32); ok {
			return augend.(float32) + result
		}
	case reflect.Float64:
		if result, ok := addend.(float64); ok {
			return augend.(float64) + result
		}
	}

	return 0
}
