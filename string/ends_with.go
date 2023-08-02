package string

import "go-dash/constants"

func endsWith(str, target string, position ...int) (bool, error) {
	pos := len(str)
	if position != nil {
		pos = position[0]
	}

	if pos > len(str) {
		return false, constants.ErrOutOfRange
	}

	if string(str[pos-1]) != target {
		return false, nil
	}

	return true, nil
}
