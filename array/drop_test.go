package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_drop_valid_int_one(t *testing.T) {
	result, err := Drop([]int{1, 2, 3}, 2)

	assert.Equal(t, []int{3}, result)
	assert.Nil(t, err)
}

func Test_drop_valid_int_two(t *testing.T) {
	result, err := Drop([]int{1, 2, 3}, 5)

	assert.Equal(t, []int{}, result)
	assert.Nil(t, err)
}

func Test_drop_valid_int32(t *testing.T) {
	result, err := Drop([]int32{1, 2, 3}, 2)

	assert.Equal(t, []int32{3}, result)
	assert.Nil(t, err)
}

func Test_drop_valid_int64(t *testing.T) {
	result, err := Drop([]int64{1, 2, 3}, 2)

	assert.Equal(t, []int64{3}, result)
	assert.Nil(t, err)
}

func Test_drop_valid_float32(t *testing.T) {
	result, err := Drop([]float32{1, 2, 3}, 2)

	assert.Equal(t, []float32{3}, result)
	assert.Nil(t, err)
}

func Test_drop_valid_float64(t *testing.T) {
	result, err := Drop([]float64{1, 2, 3}, 2)

	assert.Equal(t, []float64{3}, result)
	assert.Nil(t, err)
}

func Test_drop_valid_string(t *testing.T) {
	result, err := Drop([]string{"1", "2", "3"}, 2)

	assert.Equal(t, []string{"3"}, result)
	assert.Nil(t, err)
}

func Test_drop_valid_zero_drop(t *testing.T) {
	result, err := Drop([]int{1, 2, 3}, 0)

	assert.Equal(t, []int{1, 2, 3}, result)
	assert.Nil(t, err)
}

func Test_drop_valid_empty_drop(t *testing.T) {
	result, err := Drop([]int{1, 2, 3})

	assert.Equal(t, []int{1, 2, 3}, result)
	assert.Nil(t, err)
}

func Test_drop_valid_empty_array(t *testing.T) {
	result, err := Drop([]int{}, 1)

	assert.Equal(t, []int{}, result)
	assert.Nil(t, err)
}

func Test_drop_valid_array_interface(t *testing.T) {
	result, err := Drop([]interface{}{1, "2", true}, 2)

	assert.Equal(t, []interface{}{true}, result)
	assert.Nil(t, err)
}

func Test_drop_invalid_array_not_slice(t *testing.T) {
	result, err := Drop(true, 2)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_drop_invalid_array_not_support(t *testing.T) {
	result, err := Drop([]bool{true, false, true}, 2)

	assert.Equal(t, []bool{true}, result)
	assert.Nil(t, err)
}

func Test_drop_invalid_size_drop_less_than_zero(t *testing.T) {
	result, err := Drop([]int{1, 2, 3}, -2)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrParamLessThanZero, err)
}
