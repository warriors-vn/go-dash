package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_difference_valid_int_one(t *testing.T) {
	result, err := difference([]int{2, 1}, []int{2, 3})

	assert.Equal(t, []int{1}, result)
	assert.Nil(t, err)
}

func Test_difference_valid_int_two(t *testing.T) {
	result, err := difference([]int{2, 1})

	assert.Equal(t, []int{2, 1}, result)
	assert.Nil(t, err)
}

func Test_difference_valid_int32(t *testing.T) {
	result, err := difference([]int32{2, 1, 2, 2}, []int32{2, 3})

	assert.Equal(t, []int32{1}, result)
	assert.Nil(t, err)
}

func Test_difference_valid_int64(t *testing.T) {
	result, err := difference([]int64{2, 1, 2, 2, 5}, []int64{2, 3})

	assert.Equal(t, []int64{1, 5}, result)
	assert.Nil(t, err)
}

func Test_difference_valid_float32(t *testing.T) {
	result, err := difference([]float32{2, 1, 2, 2, 5}, []float32{2, 3})

	assert.Equal(t, []float32{1, 5}, result)
	assert.Nil(t, err)
}

func Test_difference_valid_bool(t *testing.T) {
	result, err := difference([]bool{true, false, false, true, true}, []bool{false})

	assert.Equal(t, []bool{true, true, true}, result)
	assert.Nil(t, err)
}

func Test_difference_valid_float64(t *testing.T) {
	result, err := difference([]float64{2, 1, 2, 2, 5}, []float64{2, 3})

	assert.Equal(t, []float64{1, 5}, result)
	assert.Nil(t, err)
}

func Test_difference_valid_string(t *testing.T) {
	result, err := difference([]string{"2", "1", "2", "2", "5"}, []string{"2", "3"})

	assert.Equal(t, []string{"1", "5"}, result)
	assert.Nil(t, err)
}

func Test_difference_valid_int_exclude_empty(t *testing.T) {
	result, err := difference([]int{2, 1}, []int{})

	assert.Equal(t, []int{2, 1}, result)
	assert.Nil(t, err)
}

func Test_difference_invalid_array_not_slice(t *testing.T) {
	result, err := difference(true, []int{2, 3})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_difference_invalid_exclude_not_slice(t *testing.T) {
	result, err := difference([]int{2, 3}, true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_difference_invalid_element_array_interface(t *testing.T) {
	result, err := difference([]interface{}{2, 3}, []int{2, 3})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_difference_invalid_array_exclude_incompatible(t *testing.T) {
	result, err := difference([]int{2, 3}, []bool{true, true})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_difference_invalid_not_support(t *testing.T) {
	result, err := difference([]uint{2, 3}, []uint{2, 3})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}
