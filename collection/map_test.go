package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_goDashMap_valid_int(t *testing.T) {
	result, err := goDashMap([]int{1, 2, 3, 4}, func(n int) int {
		return n * n
	})

	assert.Equal(t, []int{1, 4, 9, 16}, result)
	assert.Nil(t, err)
}

func Test_goDashMap_valid_response_func_invalid(t *testing.T) {
	result, err := goDashMap([]int{1, 2, 3, 4}, func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, []int{}, result)
	assert.Nil(t, err)
}

func Test_goDashMap_valid_string(t *testing.T) {
	result, err := goDashMap([]string{"a", "b", "c"}, func(n string) string {
		return n + n
	})

	assert.Equal(t, []string{"aa", "bb", "cc"}, result)
	assert.Nil(t, err)
}

func Test_goDashMap_valid_struct(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	result, err := goDashMap([]*User{{Name: "Kakalot", Age: 26}, {Name: "Vegeta", Age: 27}, {Name: "Trunk", Age: 10}}, func(n *User) *User {
		return &User{Name: n.Name, Age: n.Age * 2}
	})

	assert.Equal(t, []*User{{Name: "Kakalot", Age: 52}, {Name: "Vegeta", Age: 54}, {Name: "Trunk", Age: 20}}, result)
	assert.Nil(t, err)
}

func Test_goDashMap_invalid_array_not_slice(t *testing.T) {
	result, err := goDashMap(true, func(n int) int {
		return n * 2
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_goDashMap_invalid_predicate_not_func(t *testing.T) {
	result, err := goDashMap([]string{"a", "b", "c"}, true)

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotFunction, err)
}

func Test_goDashMap_invalid_param_predicate_limit(t *testing.T) {
	result, err := goDashMap([]string{"a", "b", "c"}, func(n string, m int) bool {
		return n == "a" || m%2 == 0
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_goDashMap_invalid_interface(t *testing.T) {
	result, err := goDashMap([]interface{}{"a", "b", "c"}, func(n interface{}) interface{} {
		return n
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrNotSupport, err)
}

func Test_goDashMap_invalid_incompatible(t *testing.T) {
	result, err := goDashMap([]string{"a", "b", "c"}, func(n int) int {
		return n * n
	})

	assert.Equal(t, nil, result)
	assert.Equal(t, constants.ErrIncompatible, err)
}
