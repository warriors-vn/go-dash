package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_take_valid_array_empty(t *testing.T) {
	result, err := Take([]int{}, 69)

	assert.Equal(t, []int{}, result)
	assert.Nil(t, err)
}

func Test_take_valid_int(t *testing.T) {
	result, err := Take([]int{1, 2, 3})

	assert.Equal(t, []int{1}, result)
	assert.Nil(t, err)
}

func Test_take_valid_int64(t *testing.T) {
	result, err := Take([]int64{1, 2, 3}, 2)

	assert.Equal(t, []int64{1, 2}, result)
	assert.Nil(t, err)
}

func Test_take_valid_float64(t *testing.T) {
	result, err := Take([]float64{1, 2, 3}, 5)

	assert.Equal(t, []float64{1, 2, 3}, result)
	assert.Nil(t, err)
}

func Test_take_valid_string(t *testing.T) {
	result, err := Take([]string{"1", "2", "3"}, 0)

	assert.Equal(t, []string{}, result)
	assert.Nil(t, err)
}

func Test_take_valid_bool(t *testing.T) {
	result, err := Take([]bool{true, true, true, false, true, false}, 4)

	assert.Equal(t, []bool{true, true, true, false}, result)
	assert.Nil(t, err)
}

func Test_take_valid_interface(t *testing.T) {
	result, err := Take([]interface{}{"true", true, 1, 2.2, true, false}, 4)

	assert.Equal(t, []interface{}{"true", true, 1, 2.2}, result)
	assert.Nil(t, err)
}

func Test_take_invalid_array_not_slice(t *testing.T) {
	result, err := Take(true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
