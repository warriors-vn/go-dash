package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_clamp_valid_int_one(t *testing.T) {
	result, err := clamp(-10, -5, 5)

	assert.Equal(t, -5, result)
	assert.Nil(t, err)
}

func Test_clamp_valid_int_two(t *testing.T) {
	result, err := clamp(10, -5, 5)

	assert.Equal(t, 5, result)
	assert.Nil(t, err)
}

func Test_clamp_valid_int32_one(t *testing.T) {
	result, err := clamp(int32(-10), int32(-5), int32(5))

	assert.Equal(t, int32(-5), result)
	assert.Nil(t, err)
}

func Test_clamp_valid_int32_two(t *testing.T) {
	result, err := clamp(int32(10), int32(-5), int32(5))

	assert.Equal(t, int32(5), result)
	assert.Nil(t, err)
}

func Test_clamp_valid_int64_one(t *testing.T) {
	result, err := clamp(int64(-10), int64(-5), int64(5))

	assert.Equal(t, int64(-5), result)
	assert.Nil(t, err)
}

func Test_clamp_valid_int64_two(t *testing.T) {
	result, err := clamp(int64(10), int64(-5), int64(5))

	assert.Equal(t, int64(5), result)
	assert.Nil(t, err)
}

func Test_clamp_valid_float32_one(t *testing.T) {
	result, err := clamp(float32(-10), float32(-5), float32(5))

	assert.Equal(t, float32(-5), result)
	assert.Nil(t, err)
}

func Test_clamp_valid_float32_two(t *testing.T) {
	result, err := clamp(float32(10), float32(-5), float32(5))

	assert.Equal(t, float32(5), result)
	assert.Nil(t, err)
}

func Test_clamp_valid_float64_one(t *testing.T) {
	result, err := clamp(float64(-10), float64(-5), float64(5))

	assert.Equal(t, float64(-5), result)
	assert.Nil(t, err)
}

func Test_clamp_valid_float64_two(t *testing.T) {
	result, err := clamp(float64(10), float64(-5), float64(5))

	assert.Equal(t, float64(5), result)
	assert.Nil(t, err)
}

func Test_clamp_invalid_incompatible(t *testing.T) {
	result, err := clamp(1, 2, true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_clamp_invalid_not_support(t *testing.T) {
	result, err := clamp(false, true, true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}
