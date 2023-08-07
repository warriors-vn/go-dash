package string

import (
	"strconv"
	"strings"
)

// ParseInt parses an integer from the given string 's'.
// It returns the parsed integer value and an error if parsing fails.
func ParseInt(s string) (int, error) {
	parsedInt, err := strconv.Atoi(strings.Replace(s, " ", "", -1))
	if err != nil {
		return 0, err
	}
	return parsedInt, nil
}
