package errors

// NewWithSubError creates new COM error with parent error.
func (e *OleError) NewWithSubError(hr uintptr, description string, err error) *OleError {
	return &OleError{
		hr:          hr,
		description: description,
		subError:    err,
	}
}
