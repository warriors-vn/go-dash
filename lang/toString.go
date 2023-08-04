package lang

import (
	"fmt"
	"reflect"
)

// toString converts a value to its string representation.
// The function accepts a value of any type (interface{}) and returns its string representation.
func toString(value interface{}) string {
	valueOf := reflect.ValueOf(value)

	switch valueOf.Kind() {
	case reflect.Slice, reflect.Array:
		result := ""
		for i := 0; i < valueOf.Len(); i++ {
			if i == valueOf.Len()-1 {
				result += fmt.Sprintf("%v", valueOf.Index(i).Interface())
			} else {
				result += fmt.Sprintf("%v,", valueOf.Index(i).Interface())
			}
		}

		return result
	}
	return fmt.Sprintf("%v", value)
}
