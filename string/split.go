package string

import "strings"

// split a string 's' into substrings using the specified 'separator'.
// It returns a slice of substrings. The 'limit' parameter controls the maximum
// number of substrings to return. If 'limit' is negative, there is no limit.
func split(s, separator string, limit int) []string {
	substrings := strings.Split(s, separator)
	if limit > len(substrings) {
		limit = len(substrings)
	}

	return substrings[0:limit]
}
