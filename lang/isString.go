package lang

import "reflect"

// IsString checks if a given value is a string type.
// The function accepts a value of any type (interface{}) and returns true if it is a string, false otherwise.
func IsString(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.String
}
