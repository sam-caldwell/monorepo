package counters

// Decrement - Decrement unconditionally to floor (0)
func (counter *ConditionalCounter) Decrement() (int, error) {
	defer func() {
		if counter.value > 0 {
			counter.value--
		}
	}()
	return int(counter.value), nil
}
