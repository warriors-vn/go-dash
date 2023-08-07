package collection

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_sampleSize_valid_int(t *testing.T) {
	result, err := SampleSize([]int{1, 2, 3, 4, 5, 6}, 3)

	assert.Equal(t, 3, reflect.ValueOf(result).Len())
	assert.Nil(t, err)
}

func Test_sampleSize_valid_int64(t *testing.T) {
	result, err := SampleSize([]int64{1, 2, 3, 4, 5, 6}, 2)

	assert.Equal(t, 2, reflect.ValueOf(result).Len())
	assert.Nil(t, err)
}

func Test_sampleSize_valid_float64(t *testing.T) {
	result, err := SampleSize([]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}, 9)

	assert.Equal(t, 9, reflect.ValueOf(result).Len())
	assert.Nil(t, err)
}

func Test_sampleSize_valid_string(t *testing.T) {
	result, err := SampleSize([]string{"1.1", "2.2", "3.3", "4.4", "5.5", "6.6"}, 8)

	assert.Equal(t, 8, reflect.ValueOf(result).Len())
	assert.Nil(t, err)
}

func Test_sampleSize_valid_interface(t *testing.T) {
	result, err := SampleSize([]interface{}{"1.1", 2.2, true, false, 5, 6.6}, 6)

	assert.Equal(t, 6, reflect.ValueOf(result).Len())
	assert.Nil(t, err)
}

func Test_sampleSize_invalid_array_not_slice(t *testing.T) {
	result, err := SampleSize(true, 1)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
