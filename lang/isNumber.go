package lang

// isNumber checks if a given value is a numeric type (integer or floating-point).
// The function accepts a value of any type (interface{}) and returns true if it is a numeric type, false otherwise.
func isNumber(value interface{}) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64:
		return true
	case uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64:
		return true
	default:
		return false
	}
}
