package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_math_valid_int(t *testing.T) {
	result := math(1, 9)

	assert.Equal(t, 10, result)
}

func Test_math_valid_int32(t *testing.T) {
	result := math(int32(1), int32(9))

	assert.Equal(t, int32(10), result)
}

func Test_math_valid_int64(t *testing.T) {
	result := math(int64(1), int64(9))

	assert.Equal(t, int64(10), result)
}

func Test_math_valid_float32(t *testing.T) {
	result := math(float32(1.5), float32(9))

	assert.Equal(t, float32(10.5), result)
}

func Test_math_valid_float64(t *testing.T) {
	result := math(1.5, 9.5)

	assert.Equal(t, float64(11), result)
}

func Test_math_invalid_float32(t *testing.T) {
	result := math(float32(1.5), 9)

	assert.Equal(t, 0, result)
}
