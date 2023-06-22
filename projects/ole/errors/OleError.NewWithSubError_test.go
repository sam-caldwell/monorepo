package errors

import (
	"errors"
	"testing"
)

func TestNewWithSubError(t *testing.T) {
	// Test HResult, description, and subError
	hr := uintptr(0x80004005) // Sample HRESULT
	description := "Sample COM error"
	subError := errors.New("sample sub-error")

	// Create a new OleError using NewWithSubError method
	oleErr := (&OleError{}).NewWithSubError(hr, description, subError)

	// Verify the returned OleError
	if oleErr.hr != hr {
		t.Fatalf("Expected HResult: %v, got: %v", hr, oleErr.hr)
	}
	if oleErr.description != description {
		t.Fatalf("Expected description: %q, got: %q", description, oleErr.description)
	}
	if oleErr.subError != subError {
		t.Fatalf("Expected subError: %v, got: %v", subError, oleErr.subError)
	}
}
