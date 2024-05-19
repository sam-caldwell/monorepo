package file

import (
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"testing"
)

func TestCreateTempFile(t *testing.T) {
	t.Run("Create temp file successfully", func(t *testing.T) {
		tempFile, err := CreateTempFile()

		defer func() {
			// Clean up the temporary file after the test
			if tempFile == nil {
				t.Fatal("temp file should not be nil")
			}
			if err = os.Remove(tempFile.Name()); err != nil {
				t.Fatalf("Failed to remove the temporary file: %v", err)
			}
		}()

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		// Add any additional assertions for the successful case
		// For example, you can check if the file exists, has a valid name, etc.
	})

	t.Run("Error when creating temp file", func(t *testing.T) {

		initialTmpDir := os.Getenv("TMPDIR")

		_ = os.Setenv("TMPDIR", "/path/to/invalid/temp/dir")

		defer func() {
			_ = os.Setenv("TMPDIR", initialTmpDir)
		}()
		_, err := CreateTempFile()

		if err == nil {
			t.Error("Expected an error, but got nil")
		}

		// Add any additional assertions for the error case
		// For example, you can check if the error message is as expected
	})
}
