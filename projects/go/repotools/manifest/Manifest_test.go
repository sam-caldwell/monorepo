package projectmanifest

import (
	assert2 "github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestManifest(t *testing.T) {
	manifest := Manifest{} // Create an instance of the Manifest struct

	// Test the structure of the Manifest
	assert2.Equal(t, "", manifest.fileName, "Default value of fileName should be an empty string")
	assert2.Equal(t, "", manifest.Name, "Default value of Name should be an empty string")
	assert2.Equal(t, "", manifest.Author, "Default value of Author should be an empty string")
	assert2.Equal(t, ManifestOptions{}, manifest.Options, "Default value of Options should be an empty ManifestOptions struct")
	assert2.True(t, len(manifest.SupportedPlatforms) == 0, "Default value of SupportedPlatforms should be an empty slice")
	assert2.True(t, len(manifest.Dependencies) == 0, "Default value of Dependencies should be an empty slice")
	assert2.Nil(t, manifest.err, "Default value of err should be nil")
}
