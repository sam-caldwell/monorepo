package counters

// Value - return the current value of the counter
func (counter *ConditionalCounter) Value() int {
	return int(counter.value)
}
