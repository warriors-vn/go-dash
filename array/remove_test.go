package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_remove_valid_int(t *testing.T) {
	result, old, err := Remove([]int{1, 2, 3, 4}, func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, []int{2, 4}, result)
	assert.Equal(t, []int{1, 3}, old)
	assert.Nil(t, err)
}

func Test_remove_valid_response_func_not_bool(t *testing.T) {
	result, old, err := Remove([]int{1, 2, 3, 4}, func(n int) int {
		return n
	})

	assert.Equal(t, []int{}, result)
	assert.Equal(t, []int{}, old)
	assert.Nil(t, err)
}

func Test_remove_valid_string(t *testing.T) {
	result, old, err := Remove([]string{"a", "b", "c"}, func(n string) bool {
		return n > "a"
	})

	assert.Equal(t, []string{"b", "c"}, result)
	assert.Equal(t, []string{"a"}, old)
	assert.Nil(t, err)
}

func Test_remove_valid_string_string_two(t *testing.T) {
	result, old, err := Remove([]string{"a", "b", "c"}, func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, nil, old)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_remove_valid_struct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	result, old, err := Remove([]*User{{Name: "Kakalot", Age: 26}, {Name: "Vegeta", Age: 27}, {Name: "Trunk", Age: 10}}, func(n *User) bool {
		return n.Age%2 == 0
	})

	assert.Equal(t, []*User{{Name: "Kakalot", Age: 26}, {Name: "Trunk", Age: 10}}, result)
	assert.Equal(t, []*User{{Name: "Vegeta", Age: 27}}, old)
	assert.Nil(t, err)
}

func Test_remove_invalid_array_not_slice(t *testing.T) {
	result, old, err := Remove(true, func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, nil, old)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_remove_invalid_predicate_not_func(t *testing.T) {
	result, old, err := Remove([]string{"a", "b", "c"}, true)

	assert.Equal(t, nil, result)
	assert.Equal(t, nil, old)
	assert.Equal(t, constants.ErrNotFunction, err)
}

func Test_remove_invalid_param_predicate_limit(t *testing.T) {
	result, old, err := Remove([]string{"a", "b", "c"}, func(n string, m int) bool {
		return n == "a" || m%2 == 0
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, nil, old)
	assert.Equal(t, constants.ErrNotSupport, err)
}
