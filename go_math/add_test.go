package go_math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_add_valid_int(t *testing.T) {
	result := add(1, 9)

	assert.Equal(t, 10, result)
}

func Test_add_valid_int32(t *testing.T) {
	result := add(int32(1), int32(9))

	assert.Equal(t, int32(10), result)
}

func Test_add_valid_int64(t *testing.T) {
	result := add(int64(1), int64(9))

	assert.Equal(t, int64(10), result)
}

func Test_add_valid_float32(t *testing.T) {
	result := add(float32(1.5), float32(9))

	assert.Equal(t, float32(10.5), result)
}

func Test_add_valid_float64(t *testing.T) {
	result := add(1.5, 9.5)

	assert.Equal(t, float64(11), result)
}

func Test_add_invalid_float32(t *testing.T) {
	result := add(float32(1.5), 9)

	assert.Equal(t, 0, result)
}

func Test_add_invalid_type(t *testing.T) {
	result := add(true, false)

	assert.Equal(t, 0, result)
}
