package counters

// Decrement - return the current value then decrement
func (counter *SimpleCounter) Decrement() int {
	defer func() {
		counter.value--
	}()
	return counter.value
}
