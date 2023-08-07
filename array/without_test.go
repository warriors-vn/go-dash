package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_without_valid_array_empty(t *testing.T) {
	result, err := Without([]int{}, 1, 2)

	assert.Equal(t, []int{}, result)
	assert.Nil(t, err)
}

func Test_without_valid_values_nil(t *testing.T) {
	result, err := Without([]int32{1, 2, 3})

	assert.Equal(t, []int32{1, 2, 3}, result)
	assert.Nil(t, err)
}

func Test_without_valid_values_int(t *testing.T) {
	result, err := Without([]int{2, 1, 2, 3}, 1, 2)

	assert.Equal(t, []int{3}, result)
	assert.Nil(t, err)
}

func Test_without_valid_values_int64(t *testing.T) {
	result, err := Without([]int64{2, 1, 2, 3}, int64(1), int64(2))

	assert.Equal(t, []int64{3}, result)
	assert.Nil(t, err)
}

func Test_without_valid_values_string(t *testing.T) {
	result, err := Without([]string{"2", "1", "2", "3"}, "1", 2)

	assert.Equal(t, []string{"2", "2", "3"}, result)
	assert.Nil(t, err)
}

func Test_without_valid_values_bool(t *testing.T) {
	result, err := Without([]bool{true, true, false, true, false, false}, false)

	assert.Equal(t, []bool{true, true, true}, result)
	assert.Nil(t, err)
}

func Test_without_valid_values_interface(t *testing.T) {
	result, err := Without([]interface{}{"true", true, false, 1, 2.2, 3.3}, 2.2, false)

	assert.Equal(t, []interface{}{"true", true, 1, 3.3}, result)
	assert.Nil(t, err)
}

func Test_without_invalid_array_not_slice(t *testing.T) {
	result, err := Without(true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
