package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isString_valid(t *testing.T) {
	result := isString("hihi")

	assert.Equal(t, true, result)
}

func Test_isString_invalid_one(t *testing.T) {
	result := isString(1)

	assert.Equal(t, false, result)
}

func Test_isString_invalid_two(t *testing.T) {
	result := isString(true)

	assert.Equal(t, false, result)
}
