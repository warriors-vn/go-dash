package go_math

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divide_valid(t *testing.T) {
	result := Divide(float64(6), float64(4))

	assert.Equal(t, 1.5, result)
}

func Test_divide_divisor_zero(t *testing.T) {
	result := Divide(float64(6), float64(0))

	assert.Equal(t, math.Inf(1), result)
}
