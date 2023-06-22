package errors

import (
	"testing"
)

func TestOleError_Code(t *testing.T) {
	// Create a sample OleError
	hr := uintptr(0x80004005) // Sample HRESULT
	oleErr := &OleError{
		hr:          hr,
		description: "Sample COM error",
		subError:    nil,
	}

	// Call the Code() method
	code := oleErr.Code()

	// Verify the returned HResult
	if code != hr {
		t.Errorf("Expected Code(): %v, got: %v", hr, code)
	}
}
