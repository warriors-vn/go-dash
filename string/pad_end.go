package string

// padEnd pads a string 's' with a specified 'length' by adding 'char' at the end.
// If 'char' is not provided, it defaults to a space character.
// The resulting padded string will have a total length of at least 'length'.
// If the original string length is greater than or equal to 'length', the original string is returned unchanged.
func padEnd(s string, length int, char ...string) string {
	if len(s) >= length || (char != nil && char[0] == "") {
		return s
	}

	paddingChar := " "
	if char != nil {
		paddingChar = char[0]
	}

	padLen, rightPadding, cur := length-len(s), "", 0
	for index := 0; index < padLen; index++ {
		rightPadding += string(paddingChar[cur])

		cur++
		if cur >= len(paddingChar) {
			cur = 0
		}
	}

	return s + rightPadding
}
