package projectmanifest

import (
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestSetAuthor(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct
	author := "John Doe"

	// Set the author name
	manifest.SetAuthor(author)

	// Verify that the author name is set correctly in the internal state
	assert.Equal(t, author, manifest.Author, "Author name should match the expected value")
}
