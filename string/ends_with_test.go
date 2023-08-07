package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_endsWith_valid_not_position(t *testing.T) {
	isEndWith, err := EndsWith("abc", "c")

	assert.Equal(t, true, isEndWith)
	assert.Nil(t, err)
}

func Test_endsWith_valid_position(t *testing.T) {
	isEndWith, err := EndsWith("abcdef", "c", 3)

	assert.Equal(t, true, isEndWith)
	assert.Nil(t, err)
}

func Test_endsWith_invalid_not_position(t *testing.T) {
	isEndWith, err := EndsWith("abc", "d")

	assert.Equal(t, false, isEndWith)
	assert.Nil(t, err)
}

func Test_endsWith_invalid_position(t *testing.T) {
	isEndWith, err := EndsWith("abcdef", "c", 2)

	assert.Equal(t, false, isEndWith)
	assert.Nil(t, err)
}

func Test_endsWith_out_of_range(t *testing.T) {
	isEndWith, err := EndsWith("abc", "c", 9)

	assert.Equal(t, false, isEndWith)
	assert.Equal(t, constants.ErrOutOfRange, err)
}
