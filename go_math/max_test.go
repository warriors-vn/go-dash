package go_math

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go-dash/constants"
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

func Test_maxInt32_invalid_empty(t *testing.T) {
	result := maxInt32([]int32{})

	assert.Equal(t, int32(0), result)
}

func Test_maxInt64_invalid_empty(t *testing.T) {
	result := maxInt64([]int64{})

	assert.Equal(t, int64(0), result)
}

func Test_maxFloat32_invalid_empty(t *testing.T) {
	result := maxFloat32([]float32{})

	assert.Equal(t, float32(0), result)
}

func Test_maxIntFloat64_invalid_empty(t *testing.T) {
	result := maxFloat64([]float64{})

	assert.Equal(t, float64(0), result)
}

func Test_maxInt_invalid_empty(t *testing.T) {
	result := maxInt([]int{})

	assert.Equal(t, 0, result)
}

func Test_maxField_valid(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	users := []User{
		{
			Name: "GoKu",
			Age:  26,
		},
		{
			Name: "Vegeta",
			Age:  27,
		},
		{
			Name: "Trunk",
			Age:  9,
		},
		{
			Name: "GoTen",
			Age:  8,
		},
	}

	maxAge, err := maxField(users, "Age")

	assert.Equal(t, 27, maxAge)
	assert.Nil(t, err)
}

func Test_maxField_invalid_not_slice(t *testing.T) {
	maxAge, err := maxField("not slice", "Age")

	assert.Equal(t, nil, maxAge)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_maxField_invalid_empty_slice(t *testing.T) {
	maxAge, err := maxField([]string{}, "Age")

	assert.Equal(t, nil, maxAge)
	assert.Equal(t, constants.ErrEmptyList, err)
}

func Test_maxField_invalid_field_name(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	users := []User{
		{
			Name: "Vegeta",
			Age:  27,
		},
		{
			Name: "Trunk",
			Age:  9,
		},
	}
	maxAge, err := maxField(users, "age")

	assert.Equal(t, nil, maxAge)
	assert.Equal(t, constants.ErrFieldNotFound, err)
}

func Test_maxField_invalid_incompatible(t *testing.T) {
	type User struct {
		Name string
		Age  interface{}
	}

	users := []User{
		{
			Name: "Vegeta",
			Age:  27,
		},
		{
			Name: "Trunk",
			Age:  float64(9),
		},
	}
	maxAge, err := maxField(users, "Age")

	assert.Equal(t, nil, maxAge)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_maxField_invalid_type_interface(t *testing.T) {
	type User struct {
		Name string
		Age  interface{}
	}

	users := []User{
		{
			Name: "Vegeta",
			Age:  27,
		},
		{
			Name: "Trunk",
			Age:  nil,
		},
	}
	maxAge, err := maxField(users, "Age")

	assert.Equal(t, 27, maxAge)
	assert.Nil(t, err)
}
