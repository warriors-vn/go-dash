package string

import (
	"strconv"
	"strings"
)

// parseInt parses an integer from the given string 's'.
// It returns the parsed integer value and an error if parsing fails.
func parseInt(s string) (int, error) {
	parsedInt, err := strconv.Atoi(strings.ReplaceAll(s, " ", ""))
	if err != nil {
		return 0, err
	}
	return parsedInt, nil
}
