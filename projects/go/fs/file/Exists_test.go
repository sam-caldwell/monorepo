package file

import (
	os2 "github.com/sam-caldwell/monorepo/v2/projects/go/wrappers/os"
	"testing"
)

func TestExists(t *testing.T) {
	t.Run("File exists", func(t *testing.T) {
		// Create a temporary file for testing
		tempFile, err := os2.CreateTemp("", "testfile")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer func() {
			// Clean up the temporary file after the test
			if err := os2.Remove(tempFile.Name()); err != nil {
				t.Errorf("Failed to remove the temporary file: %v", err)
			}
		}()

		exists := Exists(tempFile.Name())
		if !exists {
			t.Error("Expected the file to exist, but it doesn't")
		}
	})

	t.Run("File does not exist", func(t *testing.T) {
		// Provide a non-existing file path
		nonExistentFile := "/path/to/nonexistent/file.txt"

		exists := Exists(nonExistentFile)
		if exists {
			t.Error("Expected the file to not exist, but it does")
		}
	})
}
