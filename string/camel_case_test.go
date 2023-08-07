package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_camelCase_valid_one(t *testing.T) {
	result := CamelCase("Foo Bar")
	assert.Equal(t, "fooBar", result)
}

func Test_camelCase_valid_two(t *testing.T) {
	result := CamelCase("--foo-bar--")
	assert.Equal(t, "fooBar", result)
}

func Test_camelCase_valid_three(t *testing.T) {
	result := CamelCase("__FOO_BAR__")
	assert.Equal(t, "fooBar", result)
}

func Test_camelCase_valid_four(t *testing.T) {
	result := CamelCase("@a98@#$,.23237@#$hello-world")
	assert.Equal(t, "a9823237HelloWorld", result)
}
