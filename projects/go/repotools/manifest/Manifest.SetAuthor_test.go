package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestSetAuthor(t *testing.T) {
	manifest := &Manifest{} // Create an instance of the Manifest struct
	author := "John Doe"

	// Set the author name
	manifest.SetAuthor(author)

	// Verify that the author name is set correctly in the internal state
	assert.Equal(t, author, manifest.Author, "Author name should match the expected value")
}
