package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_castArray_valid_int(t *testing.T) {
	result, err := castArray(123)

	assert.Equal(t, []int{123}, result)
	assert.Equal(t, nil, err)
}

func Test_castArray_valid_int32(t *testing.T) {
	result, err := castArray(int32(123))

	assert.Equal(t, []int32{123}, result)
	assert.Equal(t, nil, err)
}

func Test_castArray_valid_int64(t *testing.T) {
	result, err := castArray(int64(123))

	assert.Equal(t, []int64{123}, result)
	assert.Equal(t, nil, err)
}

func Test_castArray_valid_float32(t *testing.T) {
	result, err := castArray(float32(1.23))

	assert.Equal(t, []float32{1.23}, result)
	assert.Equal(t, nil, err)
}

func Test_castArray_valid_float64(t *testing.T) {
	result, err := castArray(1.23)

	assert.Equal(t, []float64{1.23}, result)
	assert.Equal(t, nil, err)
}

func Test_castArray_valid_string(t *testing.T) {
	result, err := castArray("1.23")

	assert.Equal(t, []string{"1.23"}, result)
	assert.Equal(t, nil, err)
}

func Test_castArray_valid_bool(t *testing.T) {
	result, err := castArray(true)

	assert.Equal(t, []bool{true}, result)
	assert.Equal(t, nil, err)
}

func Test_castArray_invalid_not_support(t *testing.T) {
	result, err := castArray(uint(123))

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}
