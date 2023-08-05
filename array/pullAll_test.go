package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_pullAll_valid_int(t *testing.T) {
	result, err := pullAll([]int{1, 2, 3, 4, 3}, []int{2})

	assert.Equal(t, []int{1, 3, 4, 3}, result)
	assert.Nil(t, err)
}

func Test_pullAll_valid_bool(t *testing.T) {
	result, err := pullAll([]bool{true, true, false, true, false}, []bool{false})

	assert.Equal(t, []bool{true, true, true}, result)
	assert.Nil(t, err)
}

func Test_pullAll_valid_interface(t *testing.T) {
	result, err := pullAll([]interface{}{"true", 1, 1.1, true, false}, []bool{false})

	assert.Equal(t, []interface{}{"true", 1, 1.1, true}, result)
	assert.Nil(t, err)
}

func Test_pullAll_valid_interface_two(t *testing.T) {
	result, err := pullAll([]string{"a", "b", "c", "a", "b", "c"}, []interface{}{false})

	assert.Equal(t, []string{"a", "b", "c", "a", "b", "c"}, result)
	assert.Nil(t, err)
}

func Test_pullAll_valid_string(t *testing.T) {
	result, err := pullAll([]string{"a", "b", "c", "a", "b", "c"}, []string{"a", "c"})

	assert.Equal(t, []string{"b", "b"}, result)
	assert.Nil(t, err)
}

func Test_pullAll_valid_string_two(t *testing.T) {
	result, err := pullAll([]string{"a", "b", "c", "a", "b", "c"}, []string{})

	assert.Equal(t, []string{"a", "b", "c", "a", "b", "c"}, result)
	assert.Nil(t, err)
}

func Test_pullAll_invalid_array_not_slice(t *testing.T) {
	result, err := pullAll(true, []int{1, 2, 3})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_pullAll_invalid_remmoves_not_slice(t *testing.T) {
	result, err := pullAll([]int{1, 2, 3}, true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
