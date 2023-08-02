package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kebabCase_valid_one(t *testing.T) {
	result, err := kebabCase("Foo Bar")

	assert.Equal(t, "foo-bar", result)
	assert.Nil(t, err)
}

func Test_kebabCase_valid_two(t *testing.T) {
	result, err := kebabCase("__FOO_BAR__")

	assert.Equal(t, "foo-bar", result)
	assert.Nil(t, err)
}

func Test_kebabCase_valid_three(t *testing.T) {
	result, err := kebabCase("fooBar")

	assert.Equal(t, "foobar", result)
	assert.Nil(t, err)
}

func Test_kebabCase_valid_four(t *testing.T) {
	result, err := kebabCase(")*)_%&^sa_FOO_BAR__")

	assert.Equal(t, "sa-foo-bar", result)
	assert.Nil(t, err)
}
