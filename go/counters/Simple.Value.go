package counters

// Value - return the current value of the counter
func (counter *Simple) Value() int {
	return counter.value
}
