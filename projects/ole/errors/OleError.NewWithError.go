package errors

// NewWithError creates new error with HResult.
func (e *OleError) NewWithError(hr uintptr) *OleError {
	return &OleError{hr: hr}
}
