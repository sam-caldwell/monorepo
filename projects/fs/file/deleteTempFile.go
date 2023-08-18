package file

import (
	"os"
	"testing"
)

// DeleteTempFile - Delete the temporary file
func DeleteTempFile(t *testing.T, file *os.File) {
	if err := os.Remove(file.Name()); err != nil {
		t.Fatalf("Failed to remove file: %v", err)
	}
}
