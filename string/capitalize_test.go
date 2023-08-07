package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_capitalize_valid_one(t *testing.T) {
	result := Capitalize("FRED")
	assert.Equal(t, "Fred", result)
}

func Test_capitalize_valid_two(t *testing.T) {
	result := Capitalize("fRED")
	assert.Equal(t, "Fred", result)
}

func Test_capitalize_valid_three(t *testing.T) {
	result := Capitalize("FReD 123")
	assert.Equal(t, "Fred 123", result)
}

func Test_capitalize_invalid(t *testing.T) {
	result := Capitalize("")
	assert.Equal(t, "", result)
}
