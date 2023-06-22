package errors

import (
	"fmt"
	"testing"
)

func TestOleError_Error(t *testing.T) {
	// Create a sample OleError
	hr := uintptr(0x80004005) // Sample HRESULT
	description := "Sample COM error"
	oleErr := &OleError{
		hr:          hr,
		description: description,
		subError:    nil,
	}

	// Call the Error() method
	errMsg := oleErr.Error()

	// Verify the returned error message
	expectedErrMsg := fmt.Sprintf(" (%s)", description)
	if errMsg != expectedErrMsg {
		t.Errorf("Expected Error():\n"+
			"  expected: %q\n"+
			"       got: %q\n", expectedErrMsg, errMsg)
	}
}
