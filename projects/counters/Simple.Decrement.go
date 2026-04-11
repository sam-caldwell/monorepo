package counters

// Decrement - return the current value then decrement
func (counter *Simple) Decrement() int {
	defer func() {
		counter.value--
	}()
	return counter.value
}
