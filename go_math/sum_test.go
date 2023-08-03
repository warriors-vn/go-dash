package go_math

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go-dash/constants"
)

func Test_sum_valid_int(t *testing.T) {
	result, err := sum([]int{1, 2, 3})

	assert.Equal(t, float64(6), result)
	assert.Nil(t, err)
}

func Test_sum_valid_int8(t *testing.T) {
	result, err := sum([]int8{1, 2, 3})

	assert.Equal(t, float64(6), result)
	assert.Nil(t, err)
}

func Test_sum_valid_int16(t *testing.T) {
	result, err := sum([]int16{1, 2, 3})

	assert.Equal(t, float64(6), result)
	assert.Nil(t, err)
}

func Test_sum_valid_int32(t *testing.T) {
	result, err := sum([]int32{1, 2, 3})

	assert.Equal(t, float64(6), result)
	assert.Nil(t, err)
}

func Test_sum_valid_int64(t *testing.T) {
	result, err := sum([]int64{1, 2, 3})

	assert.Equal(t, float64(6), result)
	assert.Nil(t, err)
}

func Test_sum_valid_uint(t *testing.T) {
	result, err := sum([]uint{1, 2, 3})

	assert.Equal(t, float64(6), result)
	assert.Nil(t, err)
}

func Test_sum_valid_uint8(t *testing.T) {
	result, err := sum([]uint8{1, 2, 3})

	assert.Equal(t, float64(6), result)
	assert.Nil(t, err)
}

func Test_sum_valid_uint16(t *testing.T) {
	result, err := sum([]uint16{1, 2, 3})

	assert.Equal(t, float64(6), result)
	assert.Nil(t, err)
}

func Test_sum_valid_uint32(t *testing.T) {
	result, err := sum([]uint32{1, 2, 3})

	assert.Equal(t, float64(6), result)
	assert.Nil(t, err)
}

func Test_sum_valid_uint64(t *testing.T) {
	result, err := sum([]uint64{1, 2, 3})

	assert.Equal(t, float64(6), result)
	assert.Nil(t, err)
}

func Test_sum_valid_float32(t *testing.T) {
	result, err := sum([]float32{1, 2, 3})

	assert.Equal(t, float64(6), result)
	assert.Nil(t, err)
}

func Test_sum_valid_float64(t *testing.T) {
	result, err := sum([]float64{1, 2, 3})

	assert.Equal(t, float64(6), result)
	assert.Nil(t, err)
}

func Test_sum_invalid_arr_not_slice(t *testing.T) {
	result, err := sum("123")

	assert.Equal(t, float64(0), result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_sum_invalid_type_not_support(t *testing.T) {
	result, err := sum([]string{"1", "2"})

	assert.Equal(t, float64(0), result)
	assert.Equal(t, constants.ErrNotSupport, err)
}
