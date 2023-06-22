package errors

import (
	"errors"
	"testing"
)

func TestOleError_SubError(t *testing.T) {
	// Create a sample parent error
	parentError := errors.New("parent error")

	// Create a sample OleError with a parent error
	oleErr := &OleError{
		hr:          uintptr(0x80004005), // Sample HRESULT
		description: "Sample COM error",
		subError:    parentError,
	}

	// Call the SubError() method
	subError := oleErr.SubError()

	// Verify the returned parent error
	if subError != parentError {
		t.Errorf("Expected SubError(): %v, got: %v", parentError, subError)
	}
}
