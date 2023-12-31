package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_unescape_valid_one(t *testing.T) {
	result := Unescape("fred, barney, &amp; pebbles")

	assert.Equal(t, "fred, barney, & pebbles", result)
}

func Test_unescape_valid_two(t *testing.T) {
	result := Unescape("fred &lt; barney")

	assert.Equal(t, "fred < barney", result)
}

func Test_unescape_valid_three(t *testing.T) {
	result := Unescape("fred &gt; barney")

	assert.Equal(t, "fred > barney", result)
}
