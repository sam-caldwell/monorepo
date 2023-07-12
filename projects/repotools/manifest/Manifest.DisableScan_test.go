package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestDisableScan(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Set ScanEnabled to true
	manifest.Options.ScanEnabled = true

	// Test disabling scan
	manifest.DisableScan()

	// Verify that ScanEnabled is set to false
	assert.False(t, manifest.Options.ScanEnabled, "ScanEnabled should be set to false after calling DisableScan")
}
