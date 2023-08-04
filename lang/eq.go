package lang

import "reflect"

// eq checks if two values are equal using a deep comparison.
// The function accepts two values of any type (interface{}) and returns true if they are equal, false otherwise.
// It performs a deep comparison for slices and maps, and uses regular comparison for other types.
func eq(value, other interface{}) bool {
	return reflect.DeepEqual(value, other)
}
