package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowerFirst_valid_one(t *testing.T) {
	result := lowerFirst("Fred")
	assert.Equal(t, "fred", result)
}

func Test_lowerFirst_valid_two(t *testing.T) {
	result := lowerFirst("FRED")
	assert.Equal(t, "fRED", result)
}

func Test_lowerFirst_invalid(t *testing.T) {
	result := lowerFirst("")
	assert.Equal(t, "", result)
}
