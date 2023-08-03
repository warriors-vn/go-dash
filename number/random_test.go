package number

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_random_valid_one(t *testing.T) {
	result := random(float64(5))

	if result >= float64(0) && result <= float64(5) {
		assert.Nil(t, nil)
	}

	assert.Error(t, errors.New("the func random has a error"))
}

func Test_random_valid_two(t *testing.T) {
	result := random(float64(10), 6.9)

	if result >= 6.9 && result <= float64(10) {
		assert.Nil(t, nil)
	}

	assert.Error(t, errors.New("the func random has a error"))
}
