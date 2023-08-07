package string

import "strings"

// UpperFirst converts the first letter of a string to uppercase while leaving the rest unchanged.
func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}
