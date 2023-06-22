package errors

import (
	"testing"
)

func TestNewWithError(t *testing.T) {
	// Test HResult
	hr := uintptr(0x80004005) // Sample HRESULT

	// Create a new OleError using NewWithError method
	oleErr := (&OleError{}).NewWithError(hr)

	// Verify the returned OleError
	if oleErr.hr != hr {
		t.Errorf("Expected HResult: %v, got: %v", hr, oleErr.hr)
	}
	if oleErr.description != "" {
		t.Errorf("Expected empty description, got: %q", oleErr.description)
	}
	if oleErr.subError != nil {
		t.Errorf("Expected nil subError, got: %v", oleErr.subError)
	}
}
