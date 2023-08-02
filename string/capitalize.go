package string

import "strings"

// capitalizes the first letter of a string and leaves the rest unchanged.
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	lower := strings.ToLower(s)

	return strings.ToUpper(lower[:1]) + lower[1:]
}
