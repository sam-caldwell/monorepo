package errorQueue

// Clear - Clear the error queue state
func (err *ErrorQueue) Clear() {
	*err = []error{}
}
