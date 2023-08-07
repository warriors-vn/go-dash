package string

import "unicode"

// SnakeCase converts a string to snake_case format.
// It replaces spaces and punctuation with underscores and converts letters to lowercase.
func SnakeCase(s string) string {
	nextToSnakeCase, result := false, ""

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if nextToSnakeCase {
				result += "_" + string(unicode.ToLower(char))
				nextToSnakeCase = false
			} else {
				result += string(unicode.ToLower(char))
			}
		} else {
			if len(result) > 0 {
				nextToSnakeCase = true
			}
		}
	}

	return result
}
