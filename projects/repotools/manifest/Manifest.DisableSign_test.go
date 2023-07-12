package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestDisableSign(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Set SignEnabled to true
	manifest.Options.SignEnabled = true

	// Test disabling sign
	manifest.DisableSign()

	// Verify that SignEnabled is set to false
	assert.False(t, manifest.Options.SignEnabled, "SignEnabled should be set to false after calling DisableSign")
}
