package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_padStart_valid_one(t *testing.T) {
	result := PadStart("abc", 6)
	assert.Equal(t, "   abc", result)
}

func Test_padStart_valid_two(t *testing.T) {
	result := PadStart("abc", 6, "_-")
	assert.Equal(t, "_-_abc", result)
}

func Test_padStart_valid_three(t *testing.T) {
	result := PadStart("abc", 3)
	assert.Equal(t, "abc", result)
}

func Test_padStart_invalid_length(t *testing.T) {
	result := PadStart("abc", 2)
	assert.Equal(t, "abc", result)
}

func Test_padStart_char_empty(t *testing.T) {
	result := PadStart("abc", 6, "")
	assert.Equal(t, "abc", result)
}
