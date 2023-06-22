package errors

// Description returns the error summary, if available.
func (e *OleError) Description() string {
	return e.description
}
