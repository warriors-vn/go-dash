package number

import (
	"reflect"
	"sort"

	"github.com/warriors-vn/go-dash/constants"
)

// clamp a given number to be within the specified range [lower, upper].
// The function accepts three arguments: number, lower bound (lower), and upper bound (upper),
// all of type interface{}. It returns the clamped value within the specified range and an error if any issue occurs.
func clamp(number, lower, upper interface{}) (interface{}, error) {
	typeOfNumber, typeOfLower, typeOfUpper := reflect.TypeOf(number), reflect.TypeOf(lower), reflect.TypeOf(upper)

	if !(typeOfNumber.Kind() == typeOfLower.Kind() && typeOfNumber.Kind() == typeOfUpper.Kind()) {
		return nil, constants.ErrIncompatible
	}

	switch typeOfNumber.Kind() {
	case reflect.Int:
		resultNumber, _ := number.(int)
		resultLower, _ := lower.(int)
		resultUpper, _ := upper.(int)

		arr := []int{resultNumber, resultLower, resultUpper}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})

		return arr[1], nil
	case reflect.Int32:
		resultNumber, _ := number.(int32)
		resultLower, _ := lower.(int32)
		resultUpper, _ := upper.(int32)

		arr := []int32{resultNumber, resultLower, resultUpper}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})

		return arr[1], nil
	case reflect.Int64:
		resultNumber, _ := number.(int64)
		resultLower, _ := lower.(int64)
		resultUpper, _ := upper.(int64)

		arr := []int64{resultNumber, resultLower, resultUpper}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})

		return arr[1], nil
	case reflect.Float32:
		resultNumber, _ := number.(float32)
		resultLower, _ := lower.(float32)
		resultUpper, _ := upper.(float32)

		arr := []float32{resultNumber, resultLower, resultUpper}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})

		return arr[1], nil
	case reflect.Float64:
		resultNumber, _ := number.(float64)
		resultLower, _ := lower.(float64)
		resultUpper, _ := upper.(float64)

		arr := []float64{resultNumber, resultLower, resultUpper}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})

		return arr[1], nil
	}

	return nil, constants.ErrNotSupport
}
