package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_upperFirst(t *testing.T) {

}

func Test_upperFirst_valid_one(t *testing.T) {
	result := upperFirst("fred")
	assert.Equal(t, "Fred", result)
}

func Test_upperFirst_valid_two(t *testing.T) {
	result := upperFirst("FRED")
	assert.Equal(t, "FRED", result)
}

func Test_upperFirst_invalid(t *testing.T) {
	result := upperFirst("")
	assert.Equal(t, "", result)
}
