package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_replace_valid(t *testing.T) {
	result := replace("Hi Fred", "Fred", "Barney")

	assert.Equal(t, "Hi Barney", result)
}
