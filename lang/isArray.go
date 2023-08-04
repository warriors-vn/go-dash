package lang

import (
	"reflect"
)

// isArray checks if a given input is an array (slice).
// The function accepts a value of any type (interface{}) and returns true if it is an array, false otherwise.
func isArray(input interface{}) bool {
	inputValue := reflect.ValueOf(input)

	if inputValue.Kind() == reflect.Slice || inputValue.Kind() == reflect.Array {
		return true
	}

	return false
}
