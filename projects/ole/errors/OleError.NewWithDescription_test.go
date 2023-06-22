package errors

import (
	"testing"
)

func TestNewWithDescription(t *testing.T) {
	// Test HResult and description
	hr := uintptr(0x80004005) // Sample HRESULT
	description := "Sample COM error"

	// Create a new OleError using NewWithDescription method
	oleErr := (&OleError{}).NewWithDescription(hr, description)

	// Verify the returned OleError
	if oleErr.hr != hr {
		t.Errorf("Expected HResult: %v, got: %v", hr, oleErr.hr)
	}
	if oleErr.description != description {
		t.Errorf("Expected description: %q, got: %q", description, oleErr.description)
	}
	if oleErr.subError != nil {
		t.Errorf("Expected nil subError, got: %v", oleErr.subError)
	}
}
