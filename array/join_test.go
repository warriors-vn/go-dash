package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_join_valid_int(t *testing.T) {
	result, err := join([]int{1, 2, 3}, "-")

	assert.Equal(t, "1-2-3", result)
	assert.Nil(t, err)
}

func Test_join_valid_int32(t *testing.T) {
	result, err := join([]int32{1, 2, 3}, "-")

	assert.Equal(t, "1-2-3", result)
	assert.Nil(t, err)
}

func Test_join_valid_int64(t *testing.T) {
	result, err := join([]int64{1, 2, 3}, "-")

	assert.Equal(t, "1-2-3", result)
	assert.Nil(t, err)
}

func Test_join_valid_float32(t *testing.T) {
	result, err := join([]float32{1.1, 2.2, 3.3}, "-")

	assert.Equal(t, "1.1-2.2-3.3", result)
	assert.Nil(t, err)
}

func Test_join_valid_float64(t *testing.T) {
	result, err := join([]float64{1.1, 2.2, 3.3}, "-")

	assert.Equal(t, "1.1-2.2-3.3", result)
	assert.Nil(t, err)
}

func Test_join_valid_bool(t *testing.T) {
	result, err := join([]bool{true, false}, "-")

	assert.Equal(t, "true-false", result)
	assert.Nil(t, err)
}

func Test_join_valid_string(t *testing.T) {
	result, err := join([]string{"true", "false"}, "-")

	assert.Equal(t, "true-false", result)
	assert.Nil(t, err)
}

func Test_join_valid_interface(t *testing.T) {
	result, err := join([]interface{}{"1", false, 1.1, 2.2, 2}, "-")

	assert.Equal(t, "1-false-1.1-2.2-2", result)
	assert.Nil(t, err)
}

func Test_join_invalid_array_not_slice(t *testing.T) {
	result, err := join(true, "-")

	assert.Equal(t, "", result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
