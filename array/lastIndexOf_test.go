package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_lastIndexOf_valid_int(t *testing.T) {
	result, err := lastIndexOf([]int{1, 2, 1, 2}, 2)

	assert.Equal(t, 3, result)
	assert.Nil(t, err)
}

func Test_lastIndexOf_valid_int32(t *testing.T) {
	result, err := lastIndexOf([]int{1, 2, 1, 2}, 2, 2)

	assert.Equal(t, 1, result)
	assert.Nil(t, err)
}

func Test_lastIndexOf_valid_int64(t *testing.T) {
	result, err := lastIndexOf([]int64{2, 2, 1, 2}, int64(2), 0)

	assert.Equal(t, 0, result)
	assert.Nil(t, err)
}

func Test_lastIndexOf_valid_float32_from_great_than_array_length(t *testing.T) {
	result, err := lastIndexOf([]float32{1.1, 2.1, 1.1, 2.1}, float32(2.1), 9)

	assert.Equal(t, 3, result)
	assert.Nil(t, err)
}

func Test_lastIndexOf_valid_float64_not_found_search(t *testing.T) {
	result, err := lastIndexOf([]float64{1.1, 2.2, 1.1, 2.2}, 2.3, 4)

	assert.Equal(t, -1, result)
	assert.Nil(t, err)
}

func Test_lastIndexOf_valid_interface(t *testing.T) {
	result, err := lastIndexOf([]interface{}{"1.1", 2.2, true, false, true, int64(123)}, true)

	assert.Equal(t, 4, result)
	assert.Nil(t, err)
}

func Test_lastIndexOf_invalid_array_not_slice(t *testing.T) {
	result, err := lastIndexOf(true, 1)

	assert.Equal(t, -1, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
