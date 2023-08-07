package collection

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/warriors-vn/go-dash/constants"
)

// SampleSize randomly selects and returns a specified number of elements from the input array.
// It takes an array-like data structure and the number of elements to Sample.
// The function returns a slice of randomly chosen elements, or an error if any occurs.
func SampleSize(array interface{}, number int) (interface{}, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	result, min, max := reflect.MakeSlice(arrValue.Type(), 0, 0), 0, arrValue.Len()-1
	for i := 0; i < number; i++ {
		rand.Seed(time.Now().UnixNano())
		result = reflect.Append(result, arrValue.Index(rand.Intn(max-min+1)+min))
	}

	return result.Interface(), nil
}
