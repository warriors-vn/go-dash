package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toString_valid_one(t *testing.T) {
	result := toString(123)

	assert.Equal(t, "123", result)
}

func Test_toString_valid_two(t *testing.T) {
	result := toString(-1)

	assert.Equal(t, "-1", result)
}

func Test_toString_valid_three(t *testing.T) {
	result := toString([]int{1, 2, 3})

	assert.Equal(t, "1,2,3", result)
}

func Test_toString_valid_four(t *testing.T) {
	result := toString(true)

	assert.Equal(t, "true", result)
}
