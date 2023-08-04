package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_lte_valid_int_one(t *testing.T) {
	result, err := lte(1, 2)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lte_valid_int_two(t *testing.T) {
	result, err := lte(2, 2)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lte_valid_int_three(t *testing.T) {
	result, err := lte(3, 2)

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_lte_valid_int32_one(t *testing.T) {
	result, err := lte(int32(1), int32(2))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lte_valid_int32_two(t *testing.T) {
	result, err := lte(int32(2), int32(2))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lte_valid_int32_three(t *testing.T) {
	result, err := lte(int32(3), int32(2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_lte_valid_int64_one(t *testing.T) {
	result, err := lte(int64(1), int64(2))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lte_valid_int64_two(t *testing.T) {
	result, err := lte(int64(2), int64(2))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lte_valid_int64_three(t *testing.T) {
	result, err := lte(int64(3), int64(2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_lte_valid_float32_one(t *testing.T) {
	result, err := lte(float32(1.1), float32(2.2))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lte_valid_float32_two(t *testing.T) {
	result, err := lte(float32(2.2), float32(2.2))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lte_valid_float32_three(t *testing.T) {
	result, err := lte(float32(3.2), float32(2.2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_lte_valid_float64_one(t *testing.T) {
	result, err := lte(1.1, 2.2)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lte_valid_float64_two(t *testing.T) {
	result, err := lte(2.2, 2.2)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_lte_valid_float64_three(t *testing.T) {
	result, err := lte(3.2, 2.2)

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_lte_invalid_incompatible(t *testing.T) {
	result, err := lte(1, true)

	assert.Equal(t, false, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_lte_invalid_not_found(t *testing.T) {
	result, err := lte(true, true)

	assert.Equal(t, false, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}
