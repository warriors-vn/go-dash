package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_slice_valid_int(t *testing.T) {
	result, err := Slice([]int{1, 2, 3, 4, 5, 6}, 1, 3)

	assert.Equal(t, []int{2, 3}, result)
	assert.Nil(t, err)
}

func Test_slice_valid_float64(t *testing.T) {
	result, err := Slice([]float64{1, 2, 3, 4, 5, 6}, 1, 4)

	assert.Equal(t, []float64{2, 3, 4}, result)
	assert.Nil(t, err)
}

func Test_slice_valid_string(t *testing.T) {
	result, err := Slice([]string{"1", "2", "3", "4", "5", "6"}, 1, 4)

	assert.Equal(t, []string{"2", "3", "4"}, result)
	assert.Nil(t, err)
}

func Test_slice_valid_bool(t *testing.T) {
	result, err := Slice([]bool{true, true, true, false, true, false, true}, 1, 4)

	assert.Equal(t, []bool{true, true, false}, result)
	assert.Nil(t, err)
}

func Test_slice_valid_interface(t *testing.T) {
	result, err := Slice([]interface{}{true, "true", true, 1.1, 2, false, true}, 1, 5)

	assert.Equal(t, []interface{}{"true", true, 1.1, 2}, result)
	assert.Nil(t, err)
}

func Test_slice_invalid_array_not_slice(t *testing.T) {
	result, err := Slice(true, 0, 1)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_slice_invalid_start_less_than_zero(t *testing.T) {
	result, err := Slice([]int{1, 2, 3}, -1, 1)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrParamLessThanZero, err)
}

func Test_slice_invalid_end_less_than_zero(t *testing.T) {
	result, err := Slice([]int32{1, 2, 3}, 1, -1)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrParamLessThanZero, err)
}

func Test_slice_invalid_end_less_than_start(t *testing.T) {
	result, err := Slice([]int32{1, 2, 3}, 1, 0)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrOutOfRange, err)
}

func Test_slice_invalid_array_length_less_than_start(t *testing.T) {
	result, err := Slice([]float32{1, 2, 3}, 6, 8)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrOutOfRange, err)
}

func Test_slice_invalid_array_length_less_than_end(t *testing.T) {
	result, err := Slice([]float64{1, 2, 3}, 1, 8)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrOutOfRange, err)
}
