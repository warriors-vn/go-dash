package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_sample_valid_int(t *testing.T) {
	result, err := Sample([]int{1, 2, 3})

	assert.Nil(t, err)
	if result == 1 || result == 2 || result == 3 {
		assert.Nil(t, nil)
	}
}

func Test_sample_valid_string(t *testing.T) {
	result, err := Sample([]string{"1", "2", "3"})

	assert.Nil(t, err)
	if result == "1" || result == "2" || result == "3" {
		assert.Nil(t, nil)
	}
}

func Test_sample_invalid_array_not_slice(t *testing.T) {
	result, err := Sample(true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
