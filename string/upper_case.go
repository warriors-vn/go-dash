package string

import (
	"regexp"
	"strings"

	"github.com/warriors-vn/go-dash/constants"
)

// UpperCase converts all characters in the input string 's' to uppercase.
// It returns the modified string with all characters in uppercase.
func UpperCase(s string) (string, error) {
	regex, _ := regexp.Compile(constants.PatternNumberAndAlphabet)

	matches := regex.FindAllString(strings.ToUpper(s), -1)

	return strings.Join(matches, " "), nil
}
