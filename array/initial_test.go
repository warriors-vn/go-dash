package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_initial_valid_empty_array(t *testing.T) {
	result, err := initial([]int{})

	assert.Equal(t, []int{}, result)
	assert.Nil(t, err)
}

func Test_initial_valid_int(t *testing.T) {
	result, err := initial([]int{1, 2, 3})

	assert.Equal(t, []int{1, 2}, result)
	assert.Nil(t, err)
}

func Test_initial_valid_int32(t *testing.T) {
	result, err := initial([]int32{1, 2, 3, 4, 5, 6})

	assert.Equal(t, []int32{1, 2, 3, 4, 5}, result)
	assert.Nil(t, err)
}

func Test_initial_valid_int64(t *testing.T) {
	result, err := initial([]int64{1, 2, 3, 4, 5, 6})

	assert.Equal(t, []int64{1, 2, 3, 4, 5}, result)
	assert.Nil(t, err)
}

func Test_initial_valid_float32(t *testing.T) {
	result, err := initial([]float32{1, 2, 3, 4, 5, 6})

	assert.Equal(t, []float32{1, 2, 3, 4, 5}, result)
	assert.Nil(t, err)
}

func Test_initial_valid_float64(t *testing.T) {
	result, err := initial([]float64{1, 2, 3, 4, 5, 6})

	assert.Equal(t, []float64{1, 2, 3, 4, 5}, result)
	assert.Nil(t, err)
}

func Test_initial_valid_string(t *testing.T) {
	result, err := initial([]string{"a", "b", "c"})

	assert.Equal(t, []string{"a", "b"}, result)
	assert.Nil(t, err)
}

func Test_initial_valid_bool(t *testing.T) {
	result, err := initial([]bool{true, true, false, true})

	assert.Equal(t, []bool{true, true, false}, result)
	assert.Nil(t, err)
}

func Test_initial_valid_interface(t *testing.T) {
	result, err := initial([]interface{}{"true", 1, 1.1, true, int32(123), false, true})

	assert.Equal(t, []interface{}{"true", 1, 1.1, true, int32(123), false}, result)
	assert.Nil(t, err)
}

func Test_initial_invalid_not_slice(t *testing.T) {
	result, err := initial(true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
