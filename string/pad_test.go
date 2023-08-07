package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pad_valid_no_char(t *testing.T) {
	result := Pad("abc", 8)
	assert.Equal(t, "  abc   ", result)
}

func Test_pad_valid_one(t *testing.T) {
	result := Pad("abc", 8, "_-")
	assert.Equal(t, "_-abc_-_", result)
}

func Test_pad_valid_two(t *testing.T) {
	result := Pad("Hello", 6, "*")
	assert.Equal(t, "Hello*", result)
}

func Test_pad_valid_three(t *testing.T) {
	result := Pad("Hello", 10, "*")
	assert.Equal(t, "**Hello***", result)
}

func Test_pad_invalid_length(t *testing.T) {
	result := Pad("abc", 2)
	assert.Equal(t, "abc", result)
}

func Test_pad_char_empty(t *testing.T) {
	result := Pad("abc", 6, "")
	assert.Equal(t, "abc", result)
}
