package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_concat_valid_int(t *testing.T) {
	result, err := concat([]int{1, 2, 3}, []int{4, 5, 6, 7, 8})

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, result)
	assert.Nil(t, err)
}

func Test_concat_valid_int32(t *testing.T) {
	result, err := concat([]int32{1, 2, 3}, []int32{4, 5, 6, 7, 8})

	assert.Equal(t, []int32{1, 2, 3, 4, 5, 6, 7, 8}, result)
	assert.Nil(t, err)
}

func Test_concat_valid_int64(t *testing.T) {
	result, err := concat([]int64{1, 2, 3}, []int64{4, 5, 6, 7, 8})

	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 7, 8}, result)
	assert.Nil(t, err)
}

func Test_concat_valid_float32(t *testing.T) {
	result, err := concat([]float32{1.1, 2.2, 3.3}, []float32{4.4, 5.5, 6.6, 7.7, 8.8})

	assert.Equal(t, []float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8}, result)
	assert.Nil(t, err)
}

func Test_concat_valid_float64(t *testing.T) {
	result, err := concat([]float64{1.1, 2.2, 3.3}, []float64{4.4, 5.5, 6.6, 7.7, 8.8})

	assert.Equal(t, []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8}, result)
	assert.Nil(t, err)
}

func Test_concat_valid_string(t *testing.T) {
	result, err := concat([]string{"a", "b"}, []string{"f", "e", "d"})

	assert.Equal(t, []string{"a", "b", "f", "e", "d"}, result)
	assert.Nil(t, err)
}

func Test_concat_valid_interface(t *testing.T) {
	result, err := concat([]interface{}{1, "2", true}, []interface{}{true, false})

	assert.Equal(t, []interface{}{1, "2", true, true, false}, result)
	assert.Nil(t, err)
}

func Test_concat_valid_bool(t *testing.T) {
	result, err := concat([]bool{true}, []bool{false})

	assert.Equal(t, []bool{true, false}, result)
	assert.Nil(t, err)
}

func Test_concat_valid_incompatible(t *testing.T) {
	result, err := concat([]interface{}{1, "2", true}, []float64{1.1, 2.2})

	assert.Equal(t, []interface{}{1, "2", true, 1.1, 2.2}, result)
	assert.Nil(t, err)
}

func Test_concat_invalid_array_empty(t *testing.T) {
	result, err := concat([]int{}, []int{1})

	assert.Equal(t, []int{1}, result)
	assert.Nil(t, err)
}

func Test_concat_invalid_extend_empty(t *testing.T) {
	result, err := concat([]int{1}, []int{})

	assert.Equal(t, []int{1}, result)
	assert.Nil(t, err)
}

func Test_concat_invalid_array_not_slice(t *testing.T) {
	result, err := concat(true, []int{1})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_concat_invalid_extend_not_slice(t *testing.T) {
	result, err := concat([]int{1}, true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_concat_invalid_int_array_extend_incompatible(t *testing.T) {
	result, err := concat([]int{1}, []float64{1.1})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_concat_invalid_int32_array_extend_incompatible(t *testing.T) {
	result, err := concat([]int32{1}, []float64{1.1})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_concat_invalid_int64_array_extend_incompatible(t *testing.T) {
	result, err := concat([]int64{1}, []float64{1.1})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_concat_invalid_float32_array_extend_incompatible(t *testing.T) {
	result, err := concat([]float32{1}, []float64{1.1})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_concat_invalid_float64_array_extend_incompatible(t *testing.T) {
	result, err := concat([]float64{1.2}, []int{1})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_concat_invalid_string_array_extend_incompatible(t *testing.T) {
	result, err := concat([]string{"1"}, []int{1})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}
