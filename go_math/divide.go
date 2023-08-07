package go_math

import "math"

// Divide calculates the division of the given 'dividend' by the 'divisor'.
// It returns the result of the division as a float64.
func Divide(dividend, divisor float64) float64 {
	if divisor == 0.0 {
		return math.Inf(1)
	}

	result := dividend / divisor
	return result
}
