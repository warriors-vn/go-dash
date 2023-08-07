package number

import (
	"math/rand"
	"time"
)

// Random generates a Random float64 number within the specified range.
// The function accepts an upper bound (upper) and an optional lower bound (lower).
// It returns a Random float64 number between lower (inclusive) and upper (exclusive).
func Random(upper float64, lower ...float64) float64 {
	lowerBound := float64(0)
	if lower != nil {
		lowerBound = lower[0]
	}

	rand.Seed(time.Now().UnixNano())

	return rand.Float64()*(upper-lowerBound) + lowerBound
}
