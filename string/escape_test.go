package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_escape_valid_one(t *testing.T) {
	result := Escape("fred, barney, & pebbles")

	assert.Equal(t, "fred, barney, &amp; pebbles", result)
}

func Test_escape_valid_two(t *testing.T) {
	result := Escape("Hello <World>")

	assert.Equal(t, "Hello &lt;World&gt;", result)
}
