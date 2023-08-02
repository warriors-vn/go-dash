package string

import (
	"regexp"
	"strings"

	"go-dash/constants"
)

// kebabCase converts a string to kebab-case format, where words are separated by hyphens.
// It removes spaces, punctuation, and converts letters to lowercase.
// Returns the kebab-case string and an error if the input string is empty.
func kebabCase(s string) (string, error) {
	regex, _ := regexp.Compile(constants.PatternNumberAndAlphabet)

	matches := regex.FindAllString(strings.ToLower(s), -1)

	return strings.Join(matches, "-"), nil
}
