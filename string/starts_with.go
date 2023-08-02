package string

import "go-dash/constants"

// startsWith checks if the string 'str' starts with the target string 'target' optionally
// from the specified 'position' in 'str'.
// It returns true if 'str' starts with 'target', and false otherwise.
// If 'position' is provided, the comparison starts from that position in 'str'.
func startsWith(str, target string, position ...int) (bool, error) {
	pos := 1
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
