package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_isContains_valid_one(t *testing.T) {
	result, err := isContains([]int{1, 2, 3}, 1)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_isContains_valid_two(t *testing.T) {
	result, err := isContains([]string{"1", "2", "3"}, "1")

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_isContains_valid_three(t *testing.T) {
	result, err := isContains([]string{"1", "2", "3"}, "9")

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_isContains_invalid_array_not_slice(t *testing.T) {
	result, err := isContains(true, 123)

	assert.Equal(t, false, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
