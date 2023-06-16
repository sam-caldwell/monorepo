package counters

// Value - return the current value of the counter
func (counter *SimpleCounter) Value() int {
	return counter.value
}
