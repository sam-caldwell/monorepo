package counters

// Increment - return the current value then increment
func (counter *Simple) Increment() int {
	defer func() {
		counter.value++
	}()
	return counter.value
}
