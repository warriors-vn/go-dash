package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isArray_valid(t *testing.T) {
	result := isArray([]int{1, 2, 3})

	assert.Equal(t, true, result)
}

func Test_isArray_invalid(t *testing.T) {
	result := isArray(123)

	assert.Equal(t, false, result)
}
