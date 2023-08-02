package string

import "strings"

// repeat the given string 's' a specified number of 'times'.
// It returns the concatenated string where 's' is repeated 'times' times.
// If 'times' is zero or negative, an empty string is returned.
func repeat(s string, times int) string {
	if times <= 0 {
		return ""
	}

	return strings.Repeat(s, times)
}
