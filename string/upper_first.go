package string

import "strings"

// upperFirst converts the first letter of a string to uppercase while leaving the rest unchanged.
func upperFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}
