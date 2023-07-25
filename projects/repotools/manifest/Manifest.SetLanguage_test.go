package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestSetLanguage(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct
	language := "go"

	// Set the project language
	manifest.SetLanguage(language)

	// Verify that the project language is set correctly in the internal state
	assert.Equal(t, language, manifest.Options.Language, "Project language should match the expected value")
}
