package string

import (
	"regexp"
	"unicode"

	"go-dash/constants"
)

// camelCase converts a string to camelCase format.
// It capitalizes the first letter of each word (except the first word),
// removes spaces and punctuation, and lower cases the rest of the letters.
func camelCase(s string) string {
	nextToUpper, result := false, ""

	for _, char := range s {
		if isMatch, err := regexp.MatchString(constants.PatternNumberAndAlphabet, string(char)); isMatch && err == nil {
			if nextToUpper {
				result += string(unicode.ToUpper(char))
				nextToUpper = false
			} else {
				result += string(unicode.ToLower(char))
			}
		} else {
			if len(result) > 0 {
				nextToUpper = true
			}
		}
	}

	return result
}
