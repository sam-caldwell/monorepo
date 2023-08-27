package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestSetLanguage(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct
	language := "go"

	// Set the project language
	manifest.SetLanguage(language)

	// Verify that the project language is set correctly in the internal state
	assert.Equal(t, language, manifest.Options.Language, "Project language should match the expected value")
}
