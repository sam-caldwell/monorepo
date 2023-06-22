package errors

// NewWithDescription creates new COM error with HResult and description.
func (e *OleError) NewWithDescription(hr uintptr, description string) *OleError {
	return &OleError{hr: hr, description: description}
}
