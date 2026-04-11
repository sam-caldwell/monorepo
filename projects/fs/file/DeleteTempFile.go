package file

import (
	"os"
	"testing"
)

// DeleteTempFile - Delete the temporary file
func DeleteTempFile(t *testing.T, file *os.File) {
	if err := Delete(file.Name()); err != nil {
		t.Fatalf("Failed to remove file: %v", err)
	}
}
