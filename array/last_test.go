package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_last_empty_array(t *testing.T) {
	result := last([]int{})

	assert.Equal(t, nil, result)
}

func Test_last_array_not_slice(t *testing.T) {
	result := last(true)

	assert.Equal(t, nil, result)
}

func Test_last_valid_one(t *testing.T) {
	result := last([]bool{true, false, false})

	assert.Equal(t, false, result)
}

func Test_last_valid_two(t *testing.T) {
	result := last([]int{1, 2, 3})

	assert.Equal(t, 3, result)
}

func Test_last_valid_three(t *testing.T) {
	result := last([]string{"a", "b", "c"})

	assert.Equal(t, "c", result)
}
