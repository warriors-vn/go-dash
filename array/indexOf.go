package array

import "reflect"

// indexOf finds the index of the first occurrence of a value in an array (slice).
// The function accepts an array (slice), a value to search for, and an optional fromIndex.
// It returns the index of the first occurrence of the value in the array.
// If the value is not found, the function returns -1.
func indexOf(array, input interface{}, fromIndex ...int) interface{} {
	arrValue, inputValue := reflect.ValueOf(array), reflect.ValueOf(input)

	if (arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array) || arrValue.Len() == 0 {
		return -1
	}

	from := 0
	if fromIndex != nil {
		from = fromIndex[0]
	}

	if from < 0 {
		from = 0
	}

	if from >= arrValue.Len() {
		return -1
	}

	for i := from; i < arrValue.Len(); i++ {
		if arrValue.Index(0).Kind() != inputValue.Kind() {
			return -1
		}

		if reflect.DeepEqual(arrValue.Index(i).Interface(), input) {
			return i
		}
	}

	return -1
}
