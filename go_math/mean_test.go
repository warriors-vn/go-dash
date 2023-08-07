package go_math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mean_valid_one(t *testing.T) {
	result := Mean([]float64{3.5, 7.2, 2.8, 9.1, 5.4})

	assert.Equal(t, 5.6, result)
}

func Test_mean_valid_two(t *testing.T) {
	result := Mean([]float64{4, 2, 8, 6})

	assert.Equal(t, float64(5), result)
}

func Test_mean_invalid_empty_array(t *testing.T) {
	result := Mean([]float64{})

	assert.Equal(t, float64(0), result)
}
