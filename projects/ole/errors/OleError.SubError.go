package errors

// SubError returns the parent error, if available.
func (e *OleError) SubError() error {
	return e.subError
}
