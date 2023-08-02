package go_math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxInt_valid(t *testing.T) {
	result := maxInt([]int{12, 45, 6, 23, 9, 67, 3})

	assert.Equal(t, 67, result)
}

func Test_maxInt32(t *testing.T) {
	result := maxInt32([]int32{12, 45, 6, 23, 9, 69, 3})

	assert.Equal(t, int32(69), result)
}

func Test_maxInt64(t *testing.T) {
	result := maxInt64([]int64{12, 45, 6, 23, 9, 69, 3})

	assert.Equal(t, int64(69), result)
}

func Test_maxFloat32(t *testing.T) {
	result := maxFloat32([]float32{12.5, 45.3, 6.8, 23.7, 9.1, 67.2, 3.0})

	assert.Equal(t, float32(67.2), result)
}

func Test_maxFloat64(t *testing.T) {
	result := maxFloat64([]float64{12.5, 45.3, 6.8, 23.7, 9.1, 67.2, 3.0})

	assert.Equal(t, 67.2, result)
}
