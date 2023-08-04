package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_eq_valid_one(t *testing.T) {
	result := eq(1, 1)

	assert.Equal(t, true, result)
}

func Test_eq_valid_two(t *testing.T) {
	result := eq(1, 2)

	assert.Equal(t, false, result)
}

func Test_eq_valid_three(t *testing.T) {
	result := eq(1.1, 1.1)

	assert.Equal(t, true, result)
}

func Test_eq_valid_four(t *testing.T) {
	result := eq(true, true)

	assert.Equal(t, true, result)
}

func Test_eq_valid_five(t *testing.T) {
	result := eq("true", "true")

	assert.Equal(t, true, result)
}

func Test_eq_valid_six(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	result := eq(User{Name: "Vegeta", Age: 26}, User{Name: "Vegeta", Age: 26})

	assert.Equal(t, true, result)
}

func Test_eq_valid_seven(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	result := eq(&User{Name: "Vegeta", Age: 26}, &User{Name: "Vegeta", Age: 26})

	assert.Equal(t, true, result)
}

func Test_eq_valid_eight(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	result := eq(&User{Name: "Vegeta", Age: 27}, &User{Name: "Vegeta", Age: 26})

	assert.Equal(t, false, result)
}

func Test_eq_invalid_type(t *testing.T) {
	result := eq("true", true)

	assert.Equal(t, false, result)
}
