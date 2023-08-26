package projectmanifest

import (
	"github.com/sam-caldwell/go/v2/projects/go/testing/assert"
	"testing"
)

func TestGetAuthor(t *testing.T) {
	manifest := &Manifest{Author: "John Doe"} // Create an instance of the Manifest struct with a specific author

	// Test getting the author
	author := manifest.GetAuthor()

	// Verify that the returned author matches the expected author
	assert.Equal(t, "John Doe", author, "Returned author should match the expected author")
}
