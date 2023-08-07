package go_math

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warriors-vn/go-dash/constants"
)

func Test_minInt_valid(t *testing.T) {
	result := MinInt([]int{12, 45, 6, 23, 9, 67, 3})

	assert.Equal(t, 3, result)
}

func Test_minInt32(t *testing.T) {
	result := MinInt32([]int32{12, 45, 6, 23, 9, 69, 3})

	assert.Equal(t, int32(3), result)
}

func Test_minInt64(t *testing.T) {
	result := MinInt64([]int64{12, 45, 6, 23, 9, 69, 3})

	assert.Equal(t, int64(3), result)
}

func Test_minFloat32(t *testing.T) {
	result := MinFloat32([]float32{12.5, 45.3, 6.8, 23.7, 9.1, 67.2, 3.0})

	assert.Equal(t, float32(3.0), result)
}

func Test_minFloat64(t *testing.T) {
	result := MinFloat64([]float64{12.5, 45.3, 6.8, 23.7, 9.1, 67.2, 3.0})

	assert.Equal(t, 3.0, result)
}

func Test_minInt32_invalid_empty(t *testing.T) {
	result := MinInt32([]int32{})

	assert.Equal(t, int32(0), result)
}

func Test_minInt64_invalid_empty(t *testing.T) {
	result := MinInt64([]int64{})

	assert.Equal(t, int64(0), result)
}

func Test_minFloat32_invalid_empty(t *testing.T) {
	result := MinFloat32([]float32{})

	assert.Equal(t, float32(0), result)
}

func Test_minIntFloat64_invalid_empty(t *testing.T) {
	result := MinFloat64([]float64{})

	assert.Equal(t, float64(0), result)
}

func Test_minInt_invalid_empty(t *testing.T) {
	result := MinInt([]int{})

	assert.Equal(t, 0, result)
}

func Test_minField_valid(t *testing.T) {
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

	minAge, err := MinField(users, "Age")

	assert.Equal(t, 8, minAge)
	assert.Nil(t, err)
}

func Test_minField_invalid_not_slice(t *testing.T) {
	minAge, err := MinField("not slice", "Age")

	assert.Equal(t, nil, minAge)
	assert.Equal(t, constants.ErrNotSlice, err)
}

func Test_minField_invalid_empty_slice(t *testing.T) {
	minAge, err := MinField([]string{}, "Age")

	assert.Equal(t, nil, minAge)
	assert.Equal(t, constants.ErrEmptyList, err)
}

func Test_minField_invalid_field_name(t *testing.T) {
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
	minAge, err := MinField(users, "age")

	assert.Equal(t, nil, minAge)
	assert.Equal(t, constants.ErrFieldNotFound, err)
}

func Test_minField_invalid_incompatible(t *testing.T) {
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
	minAge, err := MinField(users, "Age")

	assert.Equal(t, nil, minAge)
	assert.Equal(t, constants.ErrIncompatible, err)
}

func Test_minField_invalid_type_interface(t *testing.T) {
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
	minAge, err := MinField(users, "Age")

	assert.Equal(t, 27, minAge)
	assert.Nil(t, err)
}
