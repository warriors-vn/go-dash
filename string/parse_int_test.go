package string

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInt_valid_one(t *testing.T) {
	result, err := parseInt("08")

	assert.Equal(t, 8, result)
	assert.Nil(t, err)
}

func Test_parseInt_valid_two(t *testing.T) {
	result, err := parseInt("   -42")

	assert.Equal(t, -42, result)
	assert.Nil(t, err)
}

func Test_parseInt_invalid(t *testing.T) {
	result, err := parseInt("abc")

	assert.Equal(t, 0, result)
	assert.Error(t, err)
}

func Test_parseInt_invalid_out_of_range_one(t *testing.T) {
	result, err := parseInt(fmt.Sprintf("%v0", math.MaxInt))

	assert.Equal(t, 0, result)
	assert.Error(t, err)
}

func Test_parseInt_invalid_out_of_range_two(t *testing.T) {
	result, err := parseInt(fmt.Sprintf("%v0", math.MinInt))

	assert.Equal(t, 0, result)
	assert.Error(t, err)
}
