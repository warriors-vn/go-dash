package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_pullAt_valid_int(t *testing.T) {
	result, err := pullAt([]int{1, 2, 3}, []int{-2})

	assert.Equal(t, []int{}, result)
	assert.Equal(t, nil, err)
}

func Test_pullAt_valid_float64(t *testing.T) {
	result, err := pullAt([]float64{1.1, 2.2, 3.3}, []int{0, 1, 2})

	assert.Equal(t, []float64{1.1, 2.2, 3.3}, result)
	assert.Equal(t, nil, err)
}

func Test_pullAt_valid_bool(t *testing.T) {
	result, err := pullAt([]bool{true, true, false, true, false}, []int{0, 1, 3, 4})

	assert.Equal(t, []bool{true, true, true, false}, result)
	assert.Equal(t, nil, err)
}

func Test_pullAt_valid_interface(t *testing.T) {
	result, err := pullAt([]interface{}{"true", 1, false, true, 1.1, "false"}, []int{0, 1, 3, 4, 5})

	assert.Equal(t, []interface{}{"true", 1, true, 1.1, "false"}, result)
	assert.Equal(t, nil, err)
}

func Test_pullAt_valid_string(t *testing.T) {
	result, err := pullAt([]string{"a", "b", "c", "d"}, []int{1, 3})

	assert.Equal(t, []string{"b", "d"}, result)
	assert.Equal(t, nil, err)
}

func Test_pullAt_invalid_array_not_slice(t *testing.T) {
	result, err := pullAt(true, []int{1, 2})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
