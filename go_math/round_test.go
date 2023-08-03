package go_math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_round_valid_one(t *testing.T) {
	result := round(4.006)

	assert.Equal(t, float64(4), result)
}

func Test_round_valid_two(t *testing.T) {
	result := round(4.006, 2)

	assert.Equal(t, 4.01, result)
}

func Test_round_valid_three(t *testing.T) {
	result := round(4060, -2)

	assert.Equal(t, float64(4100), result)
}
