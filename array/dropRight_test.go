package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_dropRight_valid_int_one(t *testing.T) {
	result, err := dropRight([]int{1, 2, 3}, 2)

	assert.Equal(t, []int{1}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_int_two(t *testing.T) {
	result, err := dropRight([]int{1, 2, 3}, 5)

	assert.Equal(t, []int{}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_int32(t *testing.T) {
	result, err := dropRight([]int32{1, 2, 3}, 2)

	assert.Equal(t, []int32{1}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_int32_two(t *testing.T) {
	result, err := dropRight([]int32{1, 2, 3}, 5)

	assert.Equal(t, []int32{}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_int64(t *testing.T) {
	result, err := dropRight([]int64{1, 2, 3}, 2)

	assert.Equal(t, []int64{1}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_int64_two(t *testing.T) {
	result, err := dropRight([]int64{1, 2, 3}, 5)

	assert.Equal(t, []int64{}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_float32(t *testing.T) {
	result, err := dropRight([]float32{1.1, 2, 3}, 2)

	assert.Equal(t, []float32{1.1}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_float32_two(t *testing.T) {
	result, err := dropRight([]float32{1, 2, 3}, 5)

	assert.Equal(t, []float32{}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_float64(t *testing.T) {
	result, err := dropRight([]float64{1.1, 2.2, 3}, 2)

	assert.Equal(t, []float64{1.1}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_float64_two(t *testing.T) {
	result, err := dropRight([]float64{1, 2, 3}, 5)

	assert.Equal(t, []float64{}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_string(t *testing.T) {
	result, err := dropRight([]string{"1", "2", "3"}, 2)

	assert.Equal(t, []string{"1"}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_string_two(t *testing.T) {
	result, err := dropRight([]string{"1", "2", "3"}, 5)

	assert.Equal(t, []string{}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_zero_dropRight(t *testing.T) {
	result, err := dropRight([]int{1, 2, 3}, 0)

	assert.Equal(t, []int{1, 2, 3}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_empty_dropRight(t *testing.T) {
	result, err := dropRight([]int{1, 2, 3})

	assert.Equal(t, []int{1, 2, 3}, result)
	assert.Nil(t, err)
}

func Test_dropRight_valid_empty_array(t *testing.T) {
	result, err := dropRight([]int{}, 1)

	assert.Equal(t, []int{}, result)
	assert.Nil(t, err)
}

func Test_dropRight_invalid_array_not_slice(t *testing.T) {
	result, err := dropRight(true, 2)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_dropRight_invalid_array_interface(t *testing.T) {
	result, err := dropRight([]interface{}{1, 2}, 2)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_dropRight_invalid_array_not_support(t *testing.T) {
	result, err := dropRight([]bool{true, false}, 2)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_dropRight_invalid_size_drop_less_than_zero(t *testing.T) {
	result, err := dropRight([]int{1, 2, 3}, -2)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrParamLessThanZero, err)
}
