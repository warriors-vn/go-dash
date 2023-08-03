package go_math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subtract_valid_int(t *testing.T) {
	result := subtract(10, 1)

	assert.Equal(t, 9, result)
}

func Test_subtract_valid_int32(t *testing.T) {
	result := subtract(int32(10), int32(10))

	assert.Equal(t, int32(0), result)
}

func Test_subtract_valid_int64(t *testing.T) {
	result := subtract(int64(1), int64(10))

	assert.Equal(t, int64(-9), result)
}

func Test_subtract_valid_float32(t *testing.T) {
	result := subtract(float32(10), float32(1))

	assert.Equal(t, float32(9), result)
}

func Test_subtract_valid_float64(t *testing.T) {
	result := subtract(float64(1), float64(10))

	assert.Equal(t, float64(-9), result)
}

func Test_subtract_invalid_float32(t *testing.T) {
	result := subtract(float32(1.5), 9)

	assert.Equal(t, 0, result)
}

func Test_subtract_invalid_type(t *testing.T) {
	result := subtract(true, false)

	assert.Equal(t, 0, result)
}
