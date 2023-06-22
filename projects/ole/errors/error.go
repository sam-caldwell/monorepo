package errors

// OleError stores COM errors.
type OleError struct {
	hr          uintptr
	description string
	subError    error
}
