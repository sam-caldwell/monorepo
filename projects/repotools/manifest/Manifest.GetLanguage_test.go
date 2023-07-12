package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestGetLanguage(t *testing.T) {
	manifest := &Manifest{
		Options: ManifestOptions{
			Language: "go",
		},
	} // Create an instance of the Manifest struct with a specific language

	// Test getting the language
	language := manifest.GetLanguage()

	// Verify that the returned language matches the expected language
	assert.Equal(t, "go", language, "Returned language should match the expected language")
}
