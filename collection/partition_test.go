package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_partition_valid_int(t *testing.T) {
	result, err := partition([]int{1, 2, 3, 4}, func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, [][]int{{2, 4}, {1, 3}}, result)
	assert.Nil(t, err)
}

func Test_partition_valid_response_func_invalid(t *testing.T) {
	result, err := partition([]int{1, 2, 3, 4}, func(n int) int {
		return n
	})

	assert.Equal(t, [][]int{}, result)
	assert.Nil(t, err)
}

func Test_partition_valid_string(t *testing.T) {
	result, err := partition([]string{"a", "b", "c"}, func(n string) bool {
		return n > "a"
	})

	assert.Equal(t, [][]string{{"b", "c"}, {"a"}}, result)
	assert.Nil(t, err)
}

func Test_partition_valid_struct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	result, err := partition([]*User{{Name: "Kakalot", Age: 26}, {Name: "Vegeta", Age: 27}, {Name: "Trunk", Age: 10}}, func(n *User) bool {
		return n.Age%2 == 0
	})

	assert.Equal(t, [][]*User{{{Name: "Kakalot", Age: 26}, {Name: "Trunk", Age: 10}}, {{Name: "Vegeta", Age: 27}}}, result)
	assert.Nil(t, err)
}

func Test_partition_invalid_array_not_slice(t *testing.T) {
	result, err := partition(true, func(n int) int {
		return n * 2
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_partition_invalid_predicate_not_func(t *testing.T) {
	result, err := partition([]string{"a", "b", "c"}, true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotFunction, err)
}

func Test_partition_invalid_param_predicate_limit(t *testing.T) {
	result, err := partition([]string{"a", "b", "c"}, func(n string, m int) bool {
		return n == "a" || m%2 == 0
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_partition_invalid_interface(t *testing.T) {
	result, err := partition([]interface{}{"a", "b", "c"}, func(n interface{}) interface{} {
		return n
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_partition_invalid_incompatible(t *testing.T) {
	result, err := partition([]string{"a", "b", "c"}, func(n int) int {
		return n * n
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}
