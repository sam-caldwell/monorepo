package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestCreateManifest(t *testing.T) {
	// Test creating a manifest with a given filename
	filename := "example.yaml"
	manifest := CreateManifest(filename)

	assert.Equal(t, filename, manifest.fileName, "Manifest filename should match the provided filename")
	assert.Equal(t, "", manifest.Name, "Default value of Name should be an empty string")
	assert.Equal(t, "", manifest.Author, "Default value of Author should be an empty string")
	assert.Equal(t, ManifestOptions{}, manifest.Options, "Default value of Options should be an empty ManifestOptions struct")
	assert.True(t, len(manifest.SupportedPlatforms) == 0, "Default value of SupportedPlatforms should be an empty slice")
	assert.True(t, len(manifest.Dependencies) == 0, "Default value of Dependencies should be an empty slice")
	assert.Nil(t, manifest.err, "Default value of err should be nil")
}
