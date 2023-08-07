package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowerCase_valid_one(t *testing.T) {
	result, err := LowerCase("Foo Bar")

	assert.Equal(t, "foo bar", result)
	assert.Nil(t, err)
}

func Test_lowerCase_valid_two(t *testing.T) {
	result, err := LowerCase("__FOO_BAR__")

	assert.Equal(t, "foo bar", result)
	assert.Nil(t, err)
}

func Test_lowerCase_valid_three(t *testing.T) {
	result, err := LowerCase("fooBar")

	assert.Equal(t, "foobar", result)
	assert.Nil(t, err)
}

func Test_lowerCase_valid_four(t *testing.T) {
	result, err := LowerCase(")*)_%&^sa_FOO_BAR__")

	assert.Equal(t, "sa foo bar", result)
	assert.Nil(t, err)
}
