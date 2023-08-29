package counters

// IntegerCounter - count by increment from a starting value and return the new value along the way.
func IntegerCounter(start, increment int) func() int {

	value := start //static counter

	count := func() int {
		value += increment
		return value
	}

	return count
}
