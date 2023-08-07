package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_startsWith_valid_not_position(t *testing.T) {
	isStartWith, err := StartsWith("abc", "a")

	assert.Equal(t, true, isStartWith)
	assert.Nil(t, err)
}

func Test_startsWith_valid_position(t *testing.T) {
	isStartWith, err := StartsWith("abcdef", "c", 3)

	assert.Equal(t, true, isStartWith)
	assert.Nil(t, err)
}

func Test_startsWith_invalid_not_position(t *testing.T) {
	isStartWith, err := StartsWith("abc", "b")

	assert.Equal(t, false, isStartWith)
	assert.Nil(t, err)
}

func Test_startsWith_invalid_position(t *testing.T) {
	isStartWith, err := StartsWith("abcdef", "c", 2)

	assert.Equal(t, false, isStartWith)
	assert.Nil(t, err)
}

func Test_startsWith_out_of_range(t *testing.T) {
	isStartWith, err := StartsWith("abc", "a", 9)

	assert.Equal(t, false, isStartWith)
	assert.Equal(t, constants.ErrOutOfRange, err)
}
