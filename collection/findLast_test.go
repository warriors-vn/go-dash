package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_findLastLast_valid_int(t *testing.T) {
	result, err := FindLast([]int{1, 2, 3, 4}, func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, 4, result)
	assert.Nil(t, err)
}

func Test_findLastLast_valid_string(t *testing.T) {
	result, err := FindLast([]string{"a", "b", "c"}, func(n string) bool {
		return n > "a"
	})

	assert.Equal(t, "c", result)
	assert.Nil(t, err)
}

func Test_findLastLast_valid_struct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	result, err := FindLast([]*User{{Name: "Kakalot", Age: 26}, {Name: "Vegeta", Age: 27}, {Name: "Trunk", Age: 10}}, func(n *User) bool {
		return n.Age%2 == 0
	})

	assert.Equal(t, &User{Name: "Trunk", Age: 10}, result)
	assert.Nil(t, err)
}

func Test_findLastLast_invalid_string_string(t *testing.T) {
	result, err := FindLast([]string{"a", "b", "c"}, func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_findLastLast_invalid_response_func_not_bool(t *testing.T) {
	result, err := FindLast([]int{1, 2, 3, 4}, func(n int) int {
		return n
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotFound, err)
}

func Test_findLastLast_invalid_array_not_slice(t *testing.T) {
	result, err := FindLast(true, func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_findLastLast_invalid_predicate_not_func(t *testing.T) {
	result, err := FindLast([]string{"a", "b", "c"}, true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotFunction, err)
}

func Test_findLastLast_invalid_param_predicate_limit(t *testing.T) {
	result, err := FindLast([]string{"a", "b", "c"}, func(n string, m int) bool {
		return n == "a" || m%2 == 0
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}
