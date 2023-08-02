package string

import "strings"

// lowerFirst converts the first letter of a string to lowercase while leaving the rest unchanged.
func lowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToLower(s[:1]) + s[1:]
}
