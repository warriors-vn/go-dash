package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_indexOf_valid_empty_array(t *testing.T) {
	result := IndexOf([]int{}, 1)

	assert.Equal(t, -1, result)
}

func Test_indexOf_valid_array_not_slice(t *testing.T) {
	result := IndexOf(true, 1)

	assert.Equal(t, -1, result)
}

func Test_indexOf_valid_from_great_than_array_length(t *testing.T) {
	result := IndexOf([]int64{1, 2, 3}, 1, 6)

	assert.Equal(t, -1, result)
}

func Test_indexOf_valid_from_less_than_zero(t *testing.T) {
	result := IndexOf([]float32{1.1, 2.2, 3.3}, float32(2.2), -6)

	assert.Equal(t, 1, result)
}

func Test_indexOf_valid_incompatible(t *testing.T) {
	result := IndexOf([]interface{}{1.1, "2.2", 3.3}, 2.2, 0)

	assert.Equal(t, -1, result)
}

func Test_indexOf_valid_not_found(t *testing.T) {
	result := IndexOf([]int32{1, 2, 3}, int32(4))

	assert.Equal(t, -1, result)
}

func Test_indexOf_valid_string(t *testing.T) {
	result := IndexOf([]string{"1", "2", "b", "3", "c", "b"}, "b", 2)

	assert.Equal(t, 2, result)
}
