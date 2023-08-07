package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_padEnd_valid_one(t *testing.T) {
	result := PadEnd("abc", 6)
	assert.Equal(t, "abc   ", result)
}

func Test_padEnd_valid_two(t *testing.T) {
	result := PadEnd("abc", 6, "_-")
	assert.Equal(t, "abc_-_", result)
}

func Test_padEnd_valid_three(t *testing.T) {
	result := PadEnd("abc", 3)
	assert.Equal(t, "abc", result)
}

func Test_padEnd_invalid_length(t *testing.T) {
	result := PadEnd("abc", 2)
	assert.Equal(t, "abc", result)
}

func Test_padEnd_char_empty(t *testing.T) {
	result := PadEnd("abc", 6, "")
	assert.Equal(t, "abc", result)
}
