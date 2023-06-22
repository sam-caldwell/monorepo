//go:build windows
// +build windows

package errors

import (
	"testing"
)

func TestErrString_For_Windows(t *testing.T) {
	// Test error code
	errno := 123

	// Call the errString function
	result := errString(errno)

	// Verify the returned result
	if result != "" {
		t.Errorf("Expected errString(%d) to return an empty string, got: %q", errno, result)
	}
}
