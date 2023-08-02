package string

import (
	"regexp"
	"strings"

	"go-dash/constants"
)

// lowerCase converts a string to kebab-case format, where words are separated by hyphens.
// It removes spaces, punctuation, and converts letters to lowercase.
// Returns the kebab-case string and an error if the input string is empty.
func lowerCase(s string) (string, error) {
	regex, err := regexp.Compile(constants.PatternNumberAndAlphabet)
	if err != nil {
		return "", err
	}

	matches := regex.FindAllString(strings.ToLower(s), -1)

	return strings.Join(matches, " "), nil
}
