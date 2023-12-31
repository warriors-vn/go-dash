package string

import (
	"unicode"
)

// CamelCase converts a string to CamelCase format.
// It capitalizes the first letter of each word (except the first word),
// removes spaces and punctuation, and lower cases the rest of the letters.
func CamelCase(s string) string {
	nextToUpper, result := false, ""

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
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
