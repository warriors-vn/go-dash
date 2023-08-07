package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_inRange_invalid_number_and_start(t *testing.T) {
	result, err := InRange(1, true, 9)

	assert.Equal(t, false, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_inRange_invalid_number_and_start_two(t *testing.T) {
	result, err := InRange(1, true)

	assert.Equal(t, false, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_inRange_invalid_number_and_end(t *testing.T) {
	result, err := InRange(1, 9, true)

	assert.Equal(t, false, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_inRange_invalid_not_support(t *testing.T) {
	result, err := InRange(true, false, true)

	assert.Equal(t, false, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_inRange_valid_int_one(t *testing.T) {
	result, err := InRange(4, 8)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_int_two(t *testing.T) {
	result, err := InRange(4, 2)

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_int_three(t *testing.T) {
	result, err := InRange(-2, 2)

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_int32_one(t *testing.T) {
	result, err := InRange(int32(-1), int32(-2), int32(9))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_int32_two(t *testing.T) {
	result, err := InRange(int32(-1), int32(2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_int32_three(t *testing.T) {
	result, err := InRange(int32(-1), int32(-3), int32(-2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_int64_one(t *testing.T) {
	result, err := InRange(int64(-1), int64(-2), int64(9))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_int64_two(t *testing.T) {
	result, err := InRange(int64(-1), int64(2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_int64_three(t *testing.T) {
	result, err := InRange(int64(-1), int64(-3), int64(-2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_float32_one(t *testing.T) {
	result, err := InRange(float32(-1), float32(-2), float32(9))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_float32_two(t *testing.T) {
	result, err := InRange(float32(-1), float32(2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_float32_three(t *testing.T) {
	result, err := InRange(float32(-1), float32(-3), float32(-2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_float64_one(t *testing.T) {
	result, err := InRange(float64(-1), float64(-2), float64(9))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_float64_two(t *testing.T) {
	result, err := InRange(float64(-1), float64(2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_inRange_valid_float64_three(t *testing.T) {
	result, err := InRange(float64(-1), float64(-3), float64(-2))

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}
