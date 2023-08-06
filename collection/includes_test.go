package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_includes_valid_int(t *testing.T) {
	result, err := includes([]int{1, 2, 3, 4, 5}, 3)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_includes_valid_int64(t *testing.T) {
	result, err := includes([]int64{1, 2, 3, 4, 5}, int64(5))

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_includes_valid_float64(t *testing.T) {
	result, err := includes([]float64{1.1, 2.2, 3.3, 4.4, 5.5}, 5.5, 2)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_includes_valid_string(t *testing.T) {
	result, err := includes([]string{"1.1", "2.2", "3.3", "4.4", "5.5"}, "5.5", 3)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_includes_valid_interface(t *testing.T) {
	result, err := includes([]interface{}{"1.1", true, 3.3, 4, false}, true)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_includes_valid_not_include(t *testing.T) {
	result, err := includes([]interface{}{"1.1", true, 3.3, false}, 3)

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_includes_valid_bool(t *testing.T) {
	result, err := includes([]bool{true, true, false}, true, 1)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_includes_valid_from_great_than_array_length(t *testing.T) {
	result, err := includes([]bool{true, true, false}, true, 10)

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_includes_invalid_array_not_slice(t *testing.T) {
	result, err := includes(true, true)

	assert.Equal(t, false, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
