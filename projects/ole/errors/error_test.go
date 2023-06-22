package errors

import (
	"errors"
	"testing"
)

func TestOleError(t *testing.T) {
	// Create a sample OleError
	hr := uintptr(0x80004005) // Sample HRESULT
	description := "Sample COM error"
	subError := errors.New("sample sub-error")
	oleErr := &OleError{
		hr:          hr,
		description: description,
		subError:    subError,
	}

	// Verify the type
	if _, ok := interface{}(oleErr).(error); !ok {
		t.Errorf("OleError does not implement the error interface")
	}

	// Verify the structure
	if oleErr.hr != hr {
		t.Errorf("Expected hr: %v, got: %v", hr, oleErr.hr)
	}
	if oleErr.description != description {
		t.Errorf("Expected description: %q, got: %q", description, oleErr.description)
	}
	if oleErr.subError != subError {
		t.Errorf("Expected subError: %v, got: %v", subError, oleErr.subError)
	}
}
