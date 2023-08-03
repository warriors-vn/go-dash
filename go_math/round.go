package go_math

import "math"

// round rounds a float64 number to the specified precision.
// If precision is not provided, it defaults to 0 (no decimal places).
func round(num float64, precision ...int) float64 {
	pre := 0
	if precision != nil {
		pre = precision[0]
	}
	scale := math.Pow(10, float64(pre))

	return math.Round(num*scale) / scale
}
