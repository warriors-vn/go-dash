package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_head_empty_array(t *testing.T) {
	result := head([]int{})

	assert.Equal(t, nil, result)
}

func Test_head_array_not_slice(t *testing.T) {
	result := head(true)

	assert.Equal(t, nil, result)
}

func Test_head_valid_one(t *testing.T) {
	result := head([]bool{true, false, true})

	assert.Equal(t, true, result)
}

func Test_head_valid_two(t *testing.T) {
	result := head([]int{1, 2, 3})

	assert.Equal(t, 1, result)
}

func Test_head_valid_three(t *testing.T) {
	result := head([]string{"a", "b", "c"})

	assert.Equal(t, "a", result)
}
