package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_reverse_valid_int(t *testing.T) {
	result, err := Reverse([]int{1, 2, 3})

	assert.Equal(t, []int{3, 2, 1}, result)
	assert.Nil(t, err)
}

func Test_reverse_valid_int64(t *testing.T) {
	result, err := Reverse([]int64{1, 2, 3})

	assert.Equal(t, []int64{3, 2, 1}, result)
	assert.Nil(t, err)
}

func Test_reverse_valid_float64(t *testing.T) {
	result, err := Reverse([]float64{1.1, 2.2, 3.3})

	assert.Equal(t, []float64{3.3, 2.2, 1.1}, result)
	assert.Nil(t, err)
}

func Test_reverse_valid_string(t *testing.T) {
	result, err := Reverse([]string{"1.1", "2.2", "3.3"})

	assert.Equal(t, []string{"3.3", "2.2", "1.1"}, result)
	assert.Nil(t, err)
}

func Test_reverse_valid_bool(t *testing.T) {
	result, err := Reverse([]bool{true, true, false, true, false})

	assert.Equal(t, []bool{false, true, false, true, true}, result)
	assert.Nil(t, err)
}

func Test_reverse_valid_interface(t *testing.T) {
	result, err := Reverse([]interface{}{"true", true, 1, 2.2, false})

	assert.Equal(t, []interface{}{false, 2.2, 1, true, "true"}, result)
	assert.Nil(t, err)
}

func Test_reverse_invalid_array_not_slice(t *testing.T) {
	result, err := Reverse(true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
