package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_lt_valid_int_one(t *testing.T) {
	result, err := lt(1, 2)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lt_valid_int_two(t *testing.T) {
	result, err := lt(2, 2)

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_lt_valid_int32_one(t *testing.T) {
	result, err := lt(int32(1), int32(2))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lt_valid_int32_two(t *testing.T) {
	result, err := lt(int32(2), int32(2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_lt_valid_int64_one(t *testing.T) {
	result, err := lt(int64(1), int64(2))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lt_valid_int64_two(t *testing.T) {
	result, err := lt(int64(2), int64(2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_lt_valid_float32_one(t *testing.T) {
	result, err := lt(float32(1.1), float32(2.2))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lt_valid_float32_two(t *testing.T) {
	result, err := lt(float32(2.2), float32(2.2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_lt_valid_float64_one(t *testing.T) {
	result, err := lt(1.1, 2.2)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lt_valid_float64_two(t *testing.T) {
	result, err := lt(2.2, 2.2)

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_lt_invalid_incompatible(t *testing.T) {
	result, err := lt(1, true)

	assert.Equal(t, false, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_lt_invalid_not_found(t *testing.T) {
	result, err := lt(true, true)

	assert.Equal(t, false, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}
