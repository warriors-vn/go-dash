package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_sortedUniq_valid_int(t *testing.T) {
	result, err := SortedUniq([]int{1, 1, 2, 3, 4, 4, 5})

	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
	assert.Nil(t, err)
}

func Test_sortedUniq_valid_int64(t *testing.T) {
	result, err := SortedUniq([]int64{1, 1, 2, 3, 4, 4, 5, 1, 2})

	assert.Equal(t, []int64{1, 2, 3, 4, 5}, result)
	assert.Nil(t, err)
}

func Test_sortedUniq_valid_string(t *testing.T) {
	result, err := SortedUniq([]string{"1", "1", "2", "3", "4", "4", "5"})

	assert.Equal(t, []string{"1", "2", "3", "4", "5"}, result)
	assert.Nil(t, err)
}

func Test_sortedUniq_invalid_array_not_slice(t *testing.T) {
	result, err := SortedUniq(true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
