package string

import (
	"regexp"
	"strings"

	"go-dash/constants"
)

// upperCase converts all characters in the input string 's' to uppercase.
// It returns the modified string with all characters in uppercase.
func upperCase(s string) (string, error) {
	regex, _ := regexp.Compile(constants.PatternNumberAndAlphabet)

	matches := regex.FindAllString(strings.ToUpper(s), -1)

	return strings.Join(matches, " "), nil
}
