package go_math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ceil_valid_one(t *testing.T) {
	result := ceil(4.006)

	assert.Equal(t, float64(5), result)
}

func Test_ceil_valid_two(t *testing.T) {
	result := ceil(6.004, 2)

	assert.Equal(t, 6.01, result)
}

func Test_ceil_valid_three(t *testing.T) {
	result := ceil(6040, -2)

	assert.Equal(t, float64(6100), result)
}
