package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_isContainsArray_valid_one(t *testing.T) {
	result, err := isContainsArray([]int{1, 2, 3}, 1)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_isContainsArray_valid_two(t *testing.T) {
	result, err := isContainsArray([]string{"1", "2", "3"}, "1")

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_isContainsArray_valid_three(t *testing.T) {
	result, err := isContainsArray([]string{"1", "2", "3"}, "9")

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_isContainsArray_valid_four(t *testing.T) {
	result, err := isContainsArray([]bool{true, true, false}, true)

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_isContainsArray_valid_five(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	result, err := isContainsArray([]*User{{Name: "Vegeta", Age: 27}, {Name: "Trunk", Age: 10}, {Name: "GoTen", Age: 9}}, &User{Name: "Trunk", Age: 10})

	assert.Equal(t, true, result)
	assert.Nil(t, err)
}

func Test_isContainsArray_valid_six(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	result, err := isContainsArray([]*User{{Name: "Vegeta", Age: 27}, {Name: "Trunk", Age: 10}, {Name: "GoTen", Age: 9}}, &User{Name: "Kakalot", Age: 26})

	assert.Equal(t, false, result)
	assert.Nil(t, err)
}

func Test_isContainsArray_invalid_array_not_slice(t *testing.T) {
	result, err := isContainsArray(true, 123)

	assert.Equal(t, false, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}
