package file

import (
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"io/ioutil"
	"testing"
)

func TestExistsp(t *testing.T) {
	t.Run("File exists", func(t *testing.T) {
		// Create a temporary file for testing
		tempFile, err := ioutil.TempFile("", "testfile")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer func() {
			// Clean up the temporary file after the test
			if err := os.Remove(tempFile.Name()); err != nil {
				t.Errorf("Failed to remove the temporary file: %v", err)
			}
		}()
		tf := tempFile.Name()
		exists := Existsp(&tf)
		if !exists {
			t.Error("Expected the file to exist, but it doesn't")
		}
	})

	t.Run("File does not exist", func(t *testing.T) {
		// Provide a non-existing file path
		nonExistentFile := "/path/to/nonexistent/file.txt"

		exists := Existsp(&nonExistentFile)
		if exists {
			t.Error("Expected the file to not exist, but it does")
		}
	})

	t.Run("File is a directory", func(t *testing.T) {
		// Create a temporary directory for testing
		tempDir, err := ioutil.TempDir("", "testdir")
		if err != nil {
			t.Fatalf("Failed to create temporary directory: %v", err)
		}
		defer func() {
			// Clean up the temporary directory after the test
			if err := os.RemoveAll(tempDir); err != nil {
				t.Errorf("Failed to remove the temporary directory: %v", err)
			}
		}()

		exists := Existsp(&tempDir)
		if exists {
			t.Error("Expected the file to not exist (directory instead), but it does")
		}
	})
}
