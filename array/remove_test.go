package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_remove_valid_int(t *testing.T) {
	result, err := remove([]int{1, 2, 3, 4}, func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, []int{2, 4}, result)
	assert.Nil(t, err)
}

func Test_remove_valid_string(t *testing.T) {
	result, err := remove([]string{"a", "b", "c"}, func(n string) bool {
		return n > "a"
	})

	assert.Equal(t, []string{"b", "c"}, result)
	assert.Nil(t, err)
}

func Test_remove_valid_string_string_two(t *testing.T) {
	result, err := remove([]string{"a", "b", "c"}, func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_remove_invalid_array_not_slice(t *testing.T) {
	result, err := remove(true, func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_remove_invalid_predicate_not_func(t *testing.T) {
	result, err := remove([]string{"a", "b", "c"}, true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotFunction, err)
}

func Test_remove_invalid_param_predicate_limit(t *testing.T) {
	result, err := remove([]string{"a", "b", "c"}, func(n string, m int) bool {
		return n == "a" || m%2 == 0
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}