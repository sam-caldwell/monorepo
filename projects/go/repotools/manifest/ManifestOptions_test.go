package projectmanifest

import (
	assert2 "github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestManifestOptions(t *testing.T) {
	options := ManifestOptions{} // Create an instance of the ManifestOptions struct

	// Test the structure of the ManifestOptions
	assert2.False(t, options.BuildEnabled, "Default value of BuildEnabled should be false")
	assert2.False(t, options.LintEnabled, "Default value of LintEnabled should be false")
	assert2.False(t, options.PackEnabled, "Default value of PackEnabled should be false")
	assert2.False(t, options.ScanEnabled, "Default value of ScanEnabled should be false")
	assert2.False(t, options.SignEnabled, "Default value of SignEnabled should be false")
	assert2.Equal(t, "", options.Language, "Default value of Language should be an empty string")
}
