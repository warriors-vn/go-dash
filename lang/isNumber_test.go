package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isNumber_valid_one(t *testing.T) {
	result := IsNumber(1)

	assert.Equal(t, true, result)
}

func Test_isNumber_valid_two(t *testing.T) {
	result := IsNumber(uint(1))

	assert.Equal(t, true, result)
}

func Test_isNumber_valid_three(t *testing.T) {
	result := IsNumber(1.1)

	assert.Equal(t, true, result)
}

func Test_isNumber_invalid(t *testing.T) {
	result := IsNumber(true)

	assert.Equal(t, false, result)
}
