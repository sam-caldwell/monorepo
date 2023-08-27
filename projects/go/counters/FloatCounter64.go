package counters

// FloatCounter64 - count by increment from a starting value and return the new value along the way.
func FloatCounter64(start, increment float64) func() float64 {

	value := start //static counter

	count := func() float64 {
		value += increment
		return value
	}

	return count
}
