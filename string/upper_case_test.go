package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_upperCase_valid_one(t *testing.T) {
	result, err := upperCase("--foo-bar")

	assert.Equal(t, "FOO BAR", result)
	assert.Nil(t, err)
}

func Test_upperCase_valid_two(t *testing.T) {
	result, err := upperCase("fooBar")

	assert.Equal(t, "FOOBAR", result)
	assert.Nil(t, err)
}

func Test_upperCase_valid_three(t *testing.T) {
	result, err := upperCase("__foo_bar__")

	assert.Equal(t, "FOO BAR", result)
	assert.Nil(t, err)
}
