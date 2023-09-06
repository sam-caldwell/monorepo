package counters

// FloatCounter32 - count by increment from a starting value and return the new value along the way.
func FloatCounter32(start, increment float32) func() float32 {

	value := start //static counter

	count := func() float32 {
		value += increment
		return value
	}

	return count
}
