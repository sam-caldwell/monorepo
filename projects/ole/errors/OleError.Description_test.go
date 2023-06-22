package errors

import (
	"testing"
)

func TestOleError_Description(t *testing.T) {
	// Create a sample OleError
	hr := uintptr(0x80004005) // Sample HRESULT
	description := "Sample COM error"
	oleErr := &OleError{
		hr:          hr,
		description: description,
		subError:    nil,
	}

	// Call the Description() method
	desc := oleErr.Description()

	// Verify the returned description
	if desc != description {
		t.Errorf("Expected Description(): %q, got: %q", description, desc)
	}
}
