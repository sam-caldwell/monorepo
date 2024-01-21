package errorQueue

// Size - Return the size of the queue
func (err *ErrorQueue) Size() int {
	return len(*err)
}
