package errors

import (
	"testing"
)

func TestString(t *testing.T) {
	// Test with description
	hr := uintptr(0x80004005) // Sample HRESULT
	description := "Sample COM error"
	oleErr := &OleError{
		hr:          hr,
		description: description,
	}

	expected := " (Sample COM error)"
	if result := oleErr.String(); result != expected {
		t.Fatalf("Expected: %q\n"+
			"got: %q", expected, result)
	}

	// Test without description
	oleErr.description = ""
	expected = ""
	if result := oleErr.String(); result != expected {
		t.Fatalf("Expected: %q, got: %q", expected, result)
	}
}
