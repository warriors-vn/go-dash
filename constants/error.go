package constants

import (
	"errors"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrNotSupport    = errors.New("not support")
	ErrOutOfRange    = errors.New("len out of range")
	ErrNotSlice      = errors.New("input is not a slice")
	ErrEmptyList     = errors.New("empty list")
	ErrFieldNotFound = errors.New("field not found")
	ErrIncompatible  = errors.New("incompatible field types")
)
