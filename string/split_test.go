package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_split_valid_one(t *testing.T) {
	result := Split("a-b-c", "-", 2)

	assert.Equal(t, []string{"a", "b"}, result)
}

func Test_split_valid_two(t *testing.T) {
	result := Split("a-b-c", "-", 4)

	assert.Equal(t, []string{"a", "b", "c"}, result)
}

func Test_split_valid_three(t *testing.T) {
	result := Split("a-b-c", "abc", 2)

	assert.Equal(t, []string{"a-b-c"}, result)
}

func Test_split_valid_four(t *testing.T) {
	result := Split("a-b-c", "c", 2)

	assert.Equal(t, []string{"a-b-", ""}, result)
}
