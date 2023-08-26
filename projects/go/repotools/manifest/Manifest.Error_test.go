package projectmanifest

import (
	"errors"
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestError(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Set an error in the manifest
	expectedError := errors.New("test error")
	manifest.err = expectedError

	// Test getting the error
	err := manifest.Error()

	// Verify that the returned error matches the expected error
	assert.Equal(t, expectedError, err, "Returned error should match the expected error")
}
