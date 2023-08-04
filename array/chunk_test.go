package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_chunk_valid_int(t *testing.T) {
	result, err := chunk([]int{1, 2, 3}, 1)

	assert.Equal(t, [][]int{{1}, {2}, {3}}, result)
	assert.Nil(t, err)
}

func Test_chunk_valid_int32(t *testing.T) {
	result, err := chunk([]int32{1, 2, 3}, 2)

	assert.Equal(t, [][]int32{{1, 2}, {3}}, result)
	assert.Nil(t, err)
}

func Test_chunk_valid_int64(t *testing.T) {
	result, err := chunk([]int64{1, 2, 3}, 2)

	assert.Equal(t, [][]int64{{1, 2}, {3}}, result)
	assert.Nil(t, err)
}

func Test_chunk_valid_float32(t *testing.T) {
	result, err := chunk([]float32{1.1, 2.2, 3.3}, 2)

	assert.Equal(t, [][]float32{{1.1, 2.2}, {3.3}}, result)
	assert.Nil(t, err)
}

func Test_chunk_valid_float64(t *testing.T) {
	result, err := chunk([]float64{1.1, 2.2, 3.3}, 2)

	assert.Equal(t, [][]float64{{1.1, 2.2}, {3.3}}, result)
	assert.Nil(t, err)
}

func Test_chunk_valid_string(t *testing.T) {
	result, err := chunk([]string{"a", "b", "c", "d"}, 3)

	assert.Equal(t, [][]string{{"a", "b", "c"}, {"d"}}, result)
	assert.Nil(t, err)
}

func Test_chunk_valid_size_great_than_array_length(t *testing.T) {
	result, err := chunk([]int{1, 2, 3}, 9)

	assert.Equal(t, [][]int{{1, 2, 3}}, result)
	assert.Nil(t, err)
}

func Test_chunk_invalid_array_empty(t *testing.T) {
	result, err := chunk([]int{}, 1)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrEmptyList, err)
}

func Test_chunk_invalid_size_zero(t *testing.T) {
	result, err := chunk([]int{1, 2, 3}, 0)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrEmptyList, err)
}

func Test_chunk_invalid_array_not_slice(t *testing.T) {
	result, err := chunk(true, 1)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_chunk_invalid_array_interface(t *testing.T) {
	result, err := chunk([]interface{}{1, 2, 3}, 1)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_chunk_invalid_array_not_support(t *testing.T) {
	result, err := chunk([]bool{true, false}, 1)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_chunk_invalid_size_less_than_zero(t *testing.T) {
	result, err := chunk([]int{1, 2, 3}, -1)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrParamLessThanZero, err)
}
