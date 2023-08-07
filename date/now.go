package date

import "time"

// Now returns the current timestamp in Unix time (seconds since January 1, 1970 UTC).
// It returns the timestamp as an int64 value.
func Now() int64 {
	return time.Now().UTC().Unix()
}
