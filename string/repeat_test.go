package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_repeat_valid_one(t *testing.T) {
	result := Repeat("*", 3)

	assert.Equal(t, "***", result)
}

func Test_repeat_valid_two(t *testing.T) {
	result := Repeat("abc", 2)

	assert.Equal(t, "abcabc", result)
}

func Test_repeat_valid_three(t *testing.T) {
	result := Repeat("abc", 0)

	assert.Equal(t, "", result)
}
