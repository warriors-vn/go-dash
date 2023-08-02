package string

// pad a string 's' with a specified 'length' by adding 'chars' on both sides.
// If 'char' is not provided, it defaults to a space character.
// The resulting padded string will have a total length of at least 'length'.
// If the original string length is greater than or equal to 'length', the original string is returned unchanged.
func pad(s string, length int, char ...string) string {
	if len(s) >= length || (char != nil && char[0] == "") {
		return s
	}

	paddingChar := " "
	if char != nil {
		paddingChar = char[0]
	}

	padLen, leftPadding, rightPadding, cur, nextCur := length-len(s), "", "", 0, false
	for index := 0; index < padLen; index++ {
		if index%2 == 0 {
			rightPadding += string(paddingChar[cur])
		} else {
			leftPadding += string(paddingChar[cur])
			nextCur = true
		}

		if nextCur {
			cur++
			nextCur = false
		}

		if cur >= len(paddingChar) {
			cur = 0
		}
	}

	return leftPadding + s + rightPadding
}
