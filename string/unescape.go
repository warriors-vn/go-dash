package string

import (
	"html"
)

// Unescape a string that contains Escape sequences and returns the unescaped string.
// The input string should be surrounded by double quotes to properly Unescape it.
// If un-escaping fails, the original input string is returned.
func Unescape(s string) string {
	return html.UnescapeString(s)
}
