package pkicore

import (
	"fmt"
	"testing"
)

func TestParsePEMBlock(t *testing.T) {
	// Happy path: successful decoding
	inputPEM := []byte("-----BEGIN EXAMPLE TYPE-----\nSGVsbG8gd29ybGQhCg==\n-----END EXAMPLE TYPE-----")
	expectedType := "EXAMPLE TYPE"
	expectedDecoded := []byte("Hello world!\n")

	result, err := ParsePEMBlock(inputPEM, expectedType)
	if err != nil {
		t.Errorf("Unexpected error in happy path: %v", err)
	}

	if string(result) != string(expectedDecoded) {
		t.Errorf("Happy path test failed. Expected decoded bytes: '%s', got: '%s'", expectedDecoded, result)
	}

	// Sad path: failed decoding (wrong PEM type)
	wrongType := "WRONG TYPE"
	_, err = ParsePEMBlock(inputPEM, wrongType)
	if err == nil {
		t.Error("Sad path test failed. Expected error, but got none.")
	} else {
		expectedErrorMsg := fmt.Sprintf("failed to decode PEM block with type %s", wrongType)
		if err.Error() != expectedErrorMsg {
			t.Errorf("Sad path test failed. Expected error message: %s, got: %v", expectedErrorMsg, err)
		}
	}

	// Sad path: invalid PEM block
	invalidPEM := []byte("invalid PEM block")
	_, err = ParsePEMBlock(invalidPEM, expectedType)
	if err == nil {
		t.Error("Sad path test failed. Expected error, but got none.")
	} else {
		expectedErrorMsg := "failed to decode PEM block with type EXAMPLE TYPE"
		if err.Error() != expectedErrorMsg {
			t.Errorf("Sad path test failed. Expected error message: %s, got: %v", expectedErrorMsg, err)
		}
	}
}
