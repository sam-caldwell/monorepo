package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestEnableScan(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct

	// Set ScanEnabled to false
	manifest.Options.ScanEnabled = false

	// Test enabling scan
	manifest.EnableScan()

	// Verify that ScanEnabled is set to true
	assert.True(t, manifest.Options.ScanEnabled, "ScanEnabled should be set to true after calling EnableScan")
}
