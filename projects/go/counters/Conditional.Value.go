package counters

// Value - return the current value of the counter
func (counter *Conditional) Value() int {
	return int(counter.value)
}
