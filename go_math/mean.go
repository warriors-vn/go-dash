package go_math

// Mean calculates the Mean (average) value of a slice of float64 numbers.
// It returns the Mean value as a float64.
func Mean(array []float64) float64 {
	if len(array) == 0 {
		return float64(0)
	}

	sumArr := float64(0)
	for _, value := range array {
		sumArr += value
	}

	return sumArr / float64(len(array))
}
