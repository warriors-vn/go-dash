package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_union_valid_arrays_int(t *testing.T) {
	result, err := Union([]int{2}, []int{1, 2})

	assert.Equal(t, []int{2, 1}, result)
	assert.Nil(t, err)
}

func Test_union_valid_arrays_int64(t *testing.T) {
	result, err := Union([]int64{2}, []int64{1, 3}, 5)

	assert.Equal(t, []int64{2, 1, 3}, result)
	assert.Nil(t, err)
}

func Test_union_valid_arrays_interface(t *testing.T) {
	result, err := Union([]interface{}{2}, []int64{1, 3}, 5, []bool{true, false})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_union_invalid_arrays_nil(t *testing.T) {
	result, err := Union()

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_union_invalid_arrays_first_not_slice(t *testing.T) {
	result, err := Union(true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_union_invalid_arrays_first_empty(t *testing.T) {
	result, err := Union([]int32{})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_union_invalid_arrays_first_not_interface(t *testing.T) {
	result, err := Union([]int32{1, 2, 3}, []bool{true})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}
