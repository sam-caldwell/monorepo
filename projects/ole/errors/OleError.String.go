package errors

// String returns the description, either manually set or formatted message with the error code.
func (e *OleError) String() string {
	if e.description != "" {
		return errString(int(e.hr)) + " (" + e.description + ")"
	}
	return errString(int(e.hr))
}
