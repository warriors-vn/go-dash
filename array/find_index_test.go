package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	Name string
	Age  int
}

func Test_findIndex_int32(t *testing.T) {
	array := []int32{1, 2, 3, 8, 7, 6, 5, 9}
	target := int32(6)

	index, found := findIndex(array, target)

	assert.Equal(t, 5, index)
	assert.Equal(t, true, found)
}

func Test_findIndex_int64(t *testing.T) {
	array := []int64{1, 2, 3, 8, 7, 6, 5, 9}
	target := int64(8)

	index, found := findIndex(array, target)

	assert.Equal(t, 3, index)
	assert.Equal(t, true, found)
}

func Test_findIndex_float32(t *testing.T) {
	array := []float32{1, 2, 3, 8, 7, 6, 5, 9}
	target := float32(3)

	index, found := findIndex(array, target)

	assert.Equal(t, 2, index)
	assert.Equal(t, true, found)
}

func Test_findIndex_string(t *testing.T) {
	array := []string{"a", "b", "c"}
	target := "b"

	index, found := findIndex(array, target)

	assert.Equal(t, 1, index)
	assert.Equal(t, true, found)
}

func Test_findIndex_struct(t *testing.T) {
	array := []*User{
		{
			Name: "warriors",
			Age:  20,
		},
		{
			Name: "god",
			Age:  10,
		},
		{
			Name: "vegeta",
			Age:  15,
		},
	}
	target := &User{
		Name: "vegeta",
		Age:  15,
	}

	index, found := findIndex(array, target)

	assert.Equal(t, 2, index)
	assert.Equal(t, true, found)
}

func Test_findIndex_invalid_array(t *testing.T) {
	array := int32(9)
	target := int32(6)

	index, found := findIndex(array, target)

	assert.Equal(t, -1, index)
	assert.Equal(t, false, found)
}

func Test_findIndex_not_found(t *testing.T) {
	array := []int32{1, 2, 3, 8, 7, 6, 5, 9}
	target := int32(10)

	index, found := findIndex(array, target)

	assert.Equal(t, -1, index)
	assert.Equal(t, false, found)
}
