package errors

// Code returns the HResult.
func (e *OleError) Code() uintptr {
	return e.hr
}
