package errors

// NewOleError creates a new OleError with the specified HResult.
func NewOleError(hr uintptr) *OleError {
	return &OleError{hr: hr}
}
