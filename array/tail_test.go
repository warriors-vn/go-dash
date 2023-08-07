package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_tail_valid_int(t *testing.T) {
	result, err := Tail([]int{1, 2, 3, 3, 2, 1})

	assert.Equal(t, []int{2, 3, 3, 2, 1}, result)
	assert.Nil(t, err)
}

func Test_tail_valid_int64(t *testing.T) {
	result, err := Tail([]int64{1, 2, 3, 3, 2, 1})

	assert.Equal(t, []int64{2, 3, 3, 2, 1}, result)
	assert.Nil(t, err)
}

func Test_tail_valid_float64(t *testing.T) {
	result, err := Tail([]float64{1.1, 2.2, 3.3, 3.3, 2.2, 1.1})

	assert.Equal(t, []float64{2.2, 3.3, 3.3, 2.2, 1.1}, result)
	assert.Nil(t, err)
}

func Test_tail_valid_string(t *testing.T) {
	result, err := Tail([]string{"1.1", "2.2", "3.3", "3.3", "2.2", "1.1"})

	assert.Equal(t, []string{"2.2", "3.3", "3.3", "2.2", "1.1"}, result)
	assert.Nil(t, err)
}

func Test_tail_valid_bool(t *testing.T) {
	result, err := Tail([]bool{false, true, true})

	assert.Equal(t, []bool{true, true}, result)
	assert.Nil(t, err)
}

func Test_tail_valid_interface(t *testing.T) {
	result, err := Tail([]interface{}{false, "true", true, 1, 2.2})

	assert.Equal(t, []interface{}{"true", true, 1, 2.2}, result)
	assert.Nil(t, err)
}

func Test_tail_valid_array_empty(t *testing.T) {
	result, err := Tail([]bool{})

	assert.Equal(t, []bool{}, result)
	assert.Nil(t, err)
}

func Test_tail_invalid_array_not_slice(t *testing.T) {
	result, err := Tail(true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
