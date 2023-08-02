package string

import (
	"html"
)

// unescape a string that contains escape sequences and returns the unescaped string.
// The input string should be surrounded by double quotes to properly unescape it.
// If un-escaping fails, the original input string is returned.
func unescape(s string) string {
	return html.UnescapeString(s)
}
