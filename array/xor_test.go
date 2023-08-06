package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_xor_valid_int(t *testing.T) {
	result, err := xor([]int{4, 4, 2, 1}, []int{2, 3, 4})

	assert.Equal(t, []int{1, 3}, result)
	assert.Nil(t, err)
}

func Test_xor_valid_int64(t *testing.T) {
	result, err := xor([]int64{2, 1}, []int64{2, 3}, []int64{2})

	assert.Equal(t, []int64{1, 3}, result)
	assert.Nil(t, err)
}

func Test_xor_valid_float64(t *testing.T) {
	result, err := xor([]float64{2, 1}, 3, []float64{2, 3}, []float64{2})

	assert.Equal(t, []float64{1, 3}, result)
	assert.Nil(t, err)
}

func Test_xor_valid_string(t *testing.T) {
	result, err := xor([]string{"2", "1"}, 3, []string{"2", "3"}, []string{"2"})

	assert.Equal(t, []string{"1", "3"}, result)
	assert.Nil(t, err)
}

func Test_xor_invalid_arrays_nil(t *testing.T) {
	result, err := xor()

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_xor_invalid_first_array_not_slice(t *testing.T) {
	result, err := xor(true, []int{1, 2, 3})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_xor_invalid_first_array_empty(t *testing.T) {
	result, err := xor([]int{}, []int{1, 2, 3})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_xor_invalid_incompatible(t *testing.T) {
	result, err := xor([]int{1, 2, 3}, []float64{1, 2, 3})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}
