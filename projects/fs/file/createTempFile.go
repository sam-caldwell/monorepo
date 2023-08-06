package file

import (
	"os"
	"testing"
)

// createTempFile - Create a temporary file for testing
func createTempFile(t *testing.T, name string, content []byte) (file *os.File) {
	file, err := os.CreateTemp("", name)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	if _, err = file.Write(content); err != nil {
		t.Fatalf("Failed to write data to the temporary file: %v", err)
	}
	return file
}
