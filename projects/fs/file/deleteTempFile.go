package file

import (
	"os"
	"testing"
)

// deleteTempFile - Delete the temporary file
func deleteTempFile(t *testing.T, file *os.File) {
	if err := os.Remove(file.Name()); err != nil {
		t.Fatalf("Failed to remove file: %v", err)
	}
}
