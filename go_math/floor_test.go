package go_math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_floor_valid_one(t *testing.T) {
	result := Floor(4.006)

	assert.Equal(t, float64(4), result)
}

func Test_floor_valid_two(t *testing.T) {
	result := Floor(0.046, 2)

	assert.Equal(t, 0.04, result)
}

func Test_floor_valid_three(t *testing.T) {
	result := Floor(4060, -2)

	assert.Equal(t, float64(4000), result)
}
