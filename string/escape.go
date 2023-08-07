package string

import "html"

// Escape converts special characters in a string to their corresponding HTML entities.
// It returns the escaped string where special characters are replaced by their entities.
func Escape(s string) string {
	return html.EscapeString(s)
}
