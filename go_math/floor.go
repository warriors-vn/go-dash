package go_math

import "math"

// floor rounds the given floating-point 'num' down to the specified 'precision'.
// The 'precision' parameter determines the number of decimal places to round to.
// The function returns the rounded number as a float64.
func floor(num float64, precision ...int) float64 {
	pre := 0
	if precision != nil {
		pre = precision[0]
	}
	scale := math.Pow(10, float64(pre))

	return math.Floor(num*scale) / scale
}
