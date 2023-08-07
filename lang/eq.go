package lang

import "reflect"

// Eq checks if two values are equal using a deep comparison.
// The function accepts two values of any type (interface{}) and returns true if they are equal, false otherwise.
// It performs a deep comparison for slices and maps, and uses regular comparison for other types.
func Eq(value, other interface{}) bool {
	return reflect.DeepEqual(value, other)
}
