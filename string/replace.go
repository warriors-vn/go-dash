package string

import "strings"

// Replace occurrences of a specified 'pattern' in the string 's' with the 'replacement' string.
// It returns the modified string where all occurrences of 'pattern' are replaced with 'replacement'.
func Replace(s, pattern, replacement string) string {
	modified := strings.Replace(s, pattern, replacement, -1)
	return modified
}
