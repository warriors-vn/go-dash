package number

import (
	"reflect"

	"go-dash/constants"
)

// inRange checks if a given number is within the specified range [start, end).
// The function accepts three arguments: number, start, and end, all of type interface{}.
// It returns a boolean indicating whether the number is within the range and an error if any issue occurs.
func inRange(number, start interface{}, end ...interface{}) (bool, error) {
	typeOfNumber, typeOfStart := reflect.TypeOf(number), reflect.TypeOf(start)

	var endRange interface{}
	if end != nil {
		endRange = end[0]
		if !(typeOfNumber.Kind() == typeOfStart.Kind() && typeOfNumber.Kind() == reflect.TypeOf(endRange).Kind()) {
			return false, constants.ErrIncompatible
		}
	} else {
		endRange, start = start, endRange
		if !(typeOfNumber.Kind() == reflect.TypeOf(endRange).Kind()) {
			return false, constants.ErrIncompatible
		}
	}

	switch typeOfNumber.Kind() {
	case reflect.Int:
		if start == nil {
			start = 0
		}

		if resultStart, ok := start.(int); ok && number.(int) <= resultStart {
			return false, nil
		}

		if resultEnd, ok := endRange.(int); ok && number.(int) >= resultEnd {
			return false, nil
		}

		return true, nil
	case reflect.Int32:
		if start == nil {
			start = int32(0)
		}

		if resultStart, ok := start.(int32); ok && number.(int32) <= resultStart {
			return false, nil
		}

		if resultEnd, ok := endRange.(int32); ok && number.(int32) >= resultEnd {
			return false, nil
		}

		return true, nil
	case reflect.Int64:
		if start == nil {
			start = int64(0)
		}

		if resultStart, ok := start.(int64); ok && number.(int64) <= resultStart {
			return false, nil
		}

		if resultEnd, ok := endRange.(int64); ok && number.(int64) >= resultEnd {
			return false, nil
		}

		return true, nil
	case reflect.Float32:
		if start == nil {
			start = float32(0)
		}

		if resultStart, ok := start.(float32); ok && number.(float32) <= resultStart {
			return false, nil
		}

		if resultEnd, ok := endRange.(float32); ok && number.(float32) >= resultEnd {
			return false, nil
		}

		return true, nil
	case reflect.Float64:
		if start == nil {
			start = float64(0)
		}

		if resultStart, ok := start.(float64); ok && number.(float64) <= resultStart {
			return false, nil
		}

		if resultEnd, ok := endRange.(float64); ok && number.(float64) >= resultEnd {
			return false, nil
		}

		return true, nil
	}

	return false, constants.ErrNotSupport
}
