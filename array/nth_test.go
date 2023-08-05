package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_nth_valid_n_great_than_array_length(t *testing.T) {
	result, err := nth([]int{1, 2, 3}, 4)

	assert.Equal(t, nil, result)
	assert.Nil(t, err)
}

func Test_nth_valid_n_less_than_zero_one(t *testing.T) {
	result, err := nth([]int32{1, 2, 3}, -3)

	assert.Equal(t, int32(1), result)
	assert.Nil(t, err)
}

func Test_nth_valid_n_less_than_zero_two(t *testing.T) {
	result, err := nth([]int32{1, 2, 3}, -4)

	assert.Equal(t, nil, result)
	assert.Nil(t, err)
}

func Test_nth_valid_int64(t *testing.T) {
	result, err := nth([]int64{1, 2, 3}, 2)

	assert.Equal(t, int64(3), result)
	assert.Nil(t, err)
}

func Test_nth_valid_float32(t *testing.T) {
	result, err := nth([]float32{1, 2, 3}, 1)

	assert.Equal(t, float32(2), result)
	assert.Nil(t, err)
}

func Test_nth_valid_float64(t *testing.T) {
	result, err := nth([]float64{1.1, 2.2, 3.3}, 2)

	assert.Equal(t, 3.3, result)
	assert.Nil(t, err)
}

func Test_nth_valid_bool(t *testing.T) {
	result, err := nth([]bool{true, true, true, false, true}, 2)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_nth_valid_string(t *testing.T) {
	result, err := nth([]string{"1.1", "2.2", "3.3"}, 2)

	assert.Equal(t, "3.3", result)
	assert.Nil(t, err)
}

func Test_nth_valid_interface(t *testing.T) {
	result, err := nth([]interface{}{"1.1", true, 2.2, float32(3.3)}, 2)

	assert.Equal(t, 2.2, result)
	assert.Nil(t, err)
}

func Test_nth_invalid_array_not_slice(t *testing.T) {
	result, err := nth(true, 1)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
