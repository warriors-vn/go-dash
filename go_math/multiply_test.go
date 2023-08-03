package go_math

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_multiply(t *testing.T) {

}

func Test_multiply_valid_int(t *testing.T) {
	result, err := multiply(10, 2)

	assert.Equal(t, 20, result)
	assert.Nil(t, err)
}

func Test_multiply_valid_int32(t *testing.T) {
	result, err := multiply(int32(10), int32(10))

	assert.Equal(t, int32(100), result)
	assert.Nil(t, err)
}

func Test_multiply_valid_int64(t *testing.T) {
	result, err := multiply(int64(-1), int64(10))

	assert.Equal(t, int64(-10), result)
	assert.Nil(t, err)
}

func Test_multiply_valid_float32(t *testing.T) {
	result, err := multiply(float32(10), float32(0.5))

	assert.Equal(t, float32(5), result)
	assert.Nil(t, err)
}

func Test_multiply_valid_float64(t *testing.T) {
	result, err := multiply(float64(1), 10.5)

	assert.Equal(t, 10.5, result)
	assert.Nil(t, err)
}

func Test_multiply_invalid_float32(t *testing.T) {
	result, err := multiply(float32(1.5), 9)

	assert.Equal(t, 0, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_multiply_invalid_type(t *testing.T) {
	result, err := multiply(true, false)

	assert.Equal(t, 0, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}
