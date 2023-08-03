package go_math

// mean calculates the mean (average) value of a slice of float64 numbers.
// It returns the mean value as a float64.
func mean(array []float64) float64 {
	if len(array) == 0 {
		return float64(0)
	}

	sumArr := float64(0)
	for _, value := range array {
		sumArr += value
	}

	return sumArr / float64(len(array))
}
