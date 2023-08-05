package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_fill_valid_int(t *testing.T) {
	result, err := fill([]int{4, 6, 8, 10}, 1, 1, 3)

	assert.Equal(t, []int{4, 1, 1, 10}, result)
	assert.Nil(t, err)
}

func Test_fill_valid_start_less_than_zero(t *testing.T) {
	result, err := fill([]int{4, 6, 8, 10}, 1, -1, 3)

	assert.Equal(t, []int{1, 1, 1, 10}, result)
	assert.Nil(t, err)
}

func Test_fill_valid_end_less_than_zero(t *testing.T) {
	result, err := fill([]int32{4, 6, 8, 10}, int32(1), -10, -3)

	assert.Equal(t, []int32{1, 6, 8, 10}, result)
	assert.Nil(t, err)
}

func Test_fill_valid_end_great_than_array_length(t *testing.T) {
	result, err := fill([]int64{4, 6, 8, 10}, int64(1), 1, 10)

	assert.Equal(t, []int64{4, 1, 1, 1}, result)
	assert.Nil(t, err)
}

func Test_fill_valid_float32(t *testing.T) {
	result, err := fill([]float32{4, 6, 8, 10}, float32(1), 2, 3)

	assert.Equal(t, []float32{4, 6, 1, 10}, result)
	assert.Nil(t, err)
}

func Test_fill_valid_float64(t *testing.T) {
	result, err := fill([]float64{4, 6, 8, 10}, float64(1), 2, 3)

	assert.Equal(t, []float64{4, 6, 1, 10}, result)
	assert.Nil(t, err)
}

func Test_fill_valid_string(t *testing.T) {
	result, err := fill([]string{"4", "6", "8", "10"}, "1", 2, 3)

	assert.Equal(t, []string{"4", "6", "1", "10"}, result)
	assert.Nil(t, err)
}

func Test_fill_valid_empty_array(t *testing.T) {
	result, err := fill([]int{}, 1, 1, 2)

	assert.Equal(t, []int{}, result)
	assert.Nil(t, err)
}

func Test_fill_valid_start_great_than_end(t *testing.T) {
	result, err := fill([]int{1, 2, 3}, 9, 3, 2)

	assert.Equal(t, []int{1, 2, 3}, result)
	assert.Nil(t, err)
}

func Test_fill_valid_array_interface(t *testing.T) {
	result, err := fill([]interface{}{1, 2, 3}, "*", 1, 2)

	assert.Equal(t, []interface{}{1, "*", 3}, result)
	assert.Nil(t, err)
}

func Test_fill_valid_bool(t *testing.T) {
	result, err := fill([]bool{true, false}, true, 1, 2)

	assert.Equal(t, []bool{true, true}, result)
	assert.Nil(t, err)
}

func Test_fill_invalid_array_not_slice(t *testing.T) {
	result, err := fill(true, 1, 1, 2)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_fill_invalid_int_incompatible(t *testing.T) {
	result, err := fill([]int{1, 2, 3}, "1", 1, 2)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_fill_invalid_int32_incompatible(t *testing.T) {
	result, err := fill([]int32{1, 2, 3}, "1", 1, 2)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_fill_invalid_int64_incompatible(t *testing.T) {
	result, err := fill([]int64{1, 2, 3}, "1", 1, 2)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_fill_invalid_float32_incompatible(t *testing.T) {
	result, err := fill([]float32{1.1, 2.2, 3.3}, "1", 1, 2)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_fill_invalid_float64_incompatible(t *testing.T) {
	result, err := fill([]float64{1.1, 2.2, 3.3}, "1", 1, 2)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_fill_invalid_string_incompatible(t *testing.T) {
	result, err := fill([]string{"1", "2", "3"}, 1, 1, 2)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}
