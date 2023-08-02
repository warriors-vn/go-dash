package go_math

// maxInt finds and returns the maximum value from the given slice of integers 'i'.
func maxInt(i []int) int {
	if len(i) == 0 {
		return 0
	}

	max := i[0]
	for _, value := range i {
		if value > max {
			max = value
		}
	}

	return max
}

// maxInt32 finds and returns the maximum value from the given slice of int32 values 'i'.
func maxInt32(i []int32) int32 {
	if len(i) == 0 {
		return 0
	}

	max := i[0]

	for _, value := range i {
		if value > max {
			max = value
		}
	}

	return max
}

// maxInt64 finds and returns the maximum value from the given slice of int64 values 'i'.
func maxInt64(i []int64) int64 {
	if len(i) == 0 {
		return 0
	}

	max := i[0]

	for _, value := range i {
		if value > max {
			max = value
		}
	}

	return max
}

// maxFloat32 finds and returns the maximum value from the given slice of float32 values 'i'.
func maxFloat32(i []float32) float32 {
	if len(i) == 0 {
		return 0
	}

	max := i[0]

	for _, value := range i {
		if value > max {
			max = value
		}
	}

	return max
}

// maxFloat64 finds and returns the maximum value from the given slice of float64 values 'i'.
func maxFloat64(i []float64) float64 {
	if len(i) == 0 {
		return 0
	}

	max := i[0]

	for _, value := range i {
		if value > max {
			max = value
		}
	}

	return max
}
