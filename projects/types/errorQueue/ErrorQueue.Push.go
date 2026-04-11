package errorQueue

// Push - Append the list of errors with the input error object.
func (err *ErrorQueue) Push(e error) {
	*err = append(*err, e)
}
