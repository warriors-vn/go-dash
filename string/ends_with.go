package string

import "github.com/warriors-vn/go-dash/constants"

// endsWith checks if the string 'str' ends with the target string 'target' optionally
// up to a specified 'position' in 'str'.
// It returns true if 'str' ends with 'target', and false otherwise.
// If 'position' is provided, only the characters up to that position in 'str' are considered.
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
