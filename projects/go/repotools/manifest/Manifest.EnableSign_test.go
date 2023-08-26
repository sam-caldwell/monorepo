package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestEnableSign(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Set SignEnabled to false
	manifest.Options.SignEnabled = false

	// Test enabling sign
	manifest.EnableSign()

	// Verify that SignEnabled is set to true
	assert.True(t, manifest.Options.SignEnabled, "SignEnabled should be set to true after calling EnableSign")
}
