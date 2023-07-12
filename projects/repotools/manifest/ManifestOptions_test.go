package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestManifestOptions(t *testing.T) {
	options := ManifestOptions{} // Create an instance of the ManifestOptions struct

	// Test the structure of the ManifestOptions
	assert.False(t, options.BuildEnabled, "Default value of BuildEnabled should be false")
	assert.False(t, options.LintEnabled, "Default value of LintEnabled should be false")
	assert.False(t, options.PackEnabled, "Default value of PackEnabled should be false")
	assert.False(t, options.ScanEnabled, "Default value of ScanEnabled should be false")
	assert.False(t, options.SignEnabled, "Default value of SignEnabled should be false")
	assert.Equal(t, "", options.Language, "Default value of Language should be an empty string")
}
