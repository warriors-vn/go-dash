package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_takeRight_valid_int(t *testing.T) {
	result, err := takeRight([]int{1, 2, 3})

	assert.Equal(t, []int{3}, result)
	assert.Nil(t, err)
}

func Test_takeRight_valid_int64(t *testing.T) {
	result, err := takeRight([]int64{1, 2, 3}, 2)

	assert.Equal(t, []int64{2, 3}, result)
	assert.Nil(t, err)
}

func Test_takeRight_valid_float64(t *testing.T) {
	result, err := takeRight([]float64{1.1, 2.2, 3.3}, 5)

	assert.Equal(t, []float64{1.1, 2.2, 3.3}, result)
	assert.Nil(t, err)
}

func Test_takeRight_valid_string(t *testing.T) {
	result, err := takeRight([]string{"1.1", "2.2", "3.3"}, 0)

	assert.Equal(t, []string{}, result)
	assert.Nil(t, err)
}

func Test_takeRight_valid_bool(t *testing.T) {
	result, err := takeRight([]bool{true, true, false, true, false}, 3)

	assert.Equal(t, []bool{false, true, false}, result)
	assert.Nil(t, err)
}

func Test_takeRight_valid_interface(t *testing.T) {
	result, err := takeRight([]interface{}{true, "true", false, "true", 1.1}, 3)

	assert.Equal(t, []interface{}{false, "true", 1.1}, result)
	assert.Nil(t, err)
}

func Test_takeRight_valid_array_empty(t *testing.T) {
	result, err := takeRight([]int{})

	assert.Equal(t, []int{}, result)
	assert.Nil(t, err)
}

func Test_takeRight_invalid_array_not_slice(t *testing.T) {
	result, err := takeRight(true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
