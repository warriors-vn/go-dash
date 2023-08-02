package string

import "html"

// escape converts special characters in a string to their corresponding HTML entities.
// It returns the escaped string where special characters are replaced by their entities.
func escape(s string) string {
	return html.EscapeString(s)
}
