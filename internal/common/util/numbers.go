package util

import "math"

// Returns the number of characters required to display `n` as a number.
func NumberLen(n int) int {
	if n == 0 {
		return 1
	}

	rem := 1
	if n < 0 {
		n = -n
		rem += 1
	}

	return int(math.Log10(float64(n))) + rem
}
