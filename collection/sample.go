package collection

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/warriors-vn/go-dash/constants"
)

// sample randomly selects and returns a single element from the input array.
// It takes an array-like data structure and returns a randomly chosen element, or an error if any occurs.
func sample(array interface{}) (interface{}, error) {
	arrValue := reflect.ValueOf(array)

	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return nil, constants.ErrNotSlice
	}

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := arrValue.Len() - 1

	return arrValue.Index(rand.Intn(max-min+1) + min).Interface(), nil
}
