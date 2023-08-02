package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_snakeCase_valid_one(t *testing.T) {
	result := snakeCase("Foo Bar")

	assert.Equal(t, "foo_bar", result)
}

func Test_snakeCase_valid_two(t *testing.T) {
	result := snakeCase("fooBar")

	assert.Equal(t, "foobar", result)
}

func Test_snakeCase_valid_three(t *testing.T) {
	result := snakeCase("--FOO-BAR--")

	assert.Equal(t, "foo_bar", result)
}

func Test_snakeCase_valid_four(t *testing.T) {
	result := snakeCase("@a98@#$,.23237@#$hello-world")

	assert.Equal(t, "a98_23237_hello_world", result)
}
