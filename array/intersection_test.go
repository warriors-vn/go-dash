package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_intersection_valid_int(t *testing.T) {
	result, err := intersection([]int{1, 2, 3}, []int{1})

	assert.Equal(t, []int{1}, result)
	assert.Nil(t, err)
}

func Test_intersection_valid_int32(t *testing.T) {
	result, err := intersection([]int32{1, 2, 3}, []int32{1, 3})

	assert.Equal(t, []int32{1, 3}, result)
	assert.Nil(t, err)
}

func Test_intersection_valid_int64(t *testing.T) {
	result, err := intersection([]int64{1, 2, 3}, []int64{1, 3})

	assert.Equal(t, []int64{1, 3}, result)
	assert.Nil(t, err)
}

func Test_intersection_valid_float32(t *testing.T) {
	result, err := intersection([]float32{1.1, 2.2, 3.3}, []float32{1.1, 3.3})

	assert.Equal(t, []float32{1.1, 3.3}, result)
	assert.Nil(t, err)
}

func Test_intersection_valid_float64(t *testing.T) {
	result, err := intersection([]float64{1.1, 2.2, 3.3}, []float64{1.1, 3.3})

	assert.Equal(t, []float64{1.1, 3.3}, result)
	assert.Nil(t, err)
}

func Test_intersection_valid_string(t *testing.T) {
	result, err := intersection([]string{"1.1", "2.2", "3.3"}, []string{"1.1", "3.3"})

	assert.Equal(t, []string{"1.1", "3.3"}, result)
	assert.Nil(t, err)
}

func Test_intersection_valid_bool(t *testing.T) {
	result, err := intersection([]bool{true, false}, []bool{true})

	assert.Equal(t, []bool{true}, result)
	assert.Nil(t, err)
}

func Test_intersection_valid_array_empty(t *testing.T) {
	result, err := intersection([]int32{}, []int32{1, 2, 3})

	assert.Equal(t, []int32{}, result)
	assert.Nil(t, err)
}

func Test_intersection_valid_other_empty(t *testing.T) {
	result, err := intersection([]int32{1, 2, 3}, []int32{})

	assert.Equal(t, []int32{}, result)
	assert.Nil(t, err)
}

func Test_intersection_invalid_array_not_slice(t *testing.T) {
	result, err := intersection(true, []int{1, 2, 3})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_intersection_invalid_other_not_slice(t *testing.T) {
	result, err := intersection([]int{1, 2, 3}, true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
