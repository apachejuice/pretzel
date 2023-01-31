package util

// Returns a substring of s limited by a and length.
func Substring(s string, a, length int) string {
	asRunes := []rune(s)
	if a >= len(asRunes) {
		return ""
	}

	if a+length > len(asRunes) {
		length = len(asRunes) - a
	}

	return string(asRunes[a : a+length])
}
