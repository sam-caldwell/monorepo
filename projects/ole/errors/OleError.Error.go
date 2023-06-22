package errors

// Error implements the error interface.
func (e *OleError) Error() string {
	return e.String()
}
