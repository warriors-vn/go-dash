package go_math

import (
	"reflect"

	"github.com/warriors-vn/go-dash/constants"
)

// sum calculates the sum of elements in a slice of numeric type.
// It returns the sum as a float64.
func sum(array interface{}) (float64, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return float64(0), constants.ErrNotSlice
	}

	sumArr := float64(0)
	for i := 0; i < arrValue.Len(); i++ {
		elem := arrValue.Index(i)
		switch elem.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			sumArr += float64(elem.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			sumArr += float64(elem.Uint())
		case reflect.Float32, reflect.Float64:
			sumArr += elem.Float()
		default:
			return float64(0), constants.ErrNotSupport
		}
	}

	return sumArr, nil
}
