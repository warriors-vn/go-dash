package constants

import (
	"errors"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrNotSupport = errors.New("not support")
	ErrOutOfRange = errors.New("len out of range")
)
