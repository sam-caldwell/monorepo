package file

import (
	"os"
	"testing"
)

func TestGetFileSize(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer func() { _ = tempFile.Close() }()
	defer func() { _ = os.Remove(tempFile.Name()) }()

	// Write some content to the temporary file
	content := "Hello, world!"
	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write content to temporary file: %v", err)
	}

	// Get the file size using the function
	fileSize := GetFileSize(tempFile)

	// Verify the file size
	expectedSize := int64(len(content))
	if fileSize != expectedSize {
		t.Fatalf("Expected file size: %d, but got: %d", expectedSize, fileSize)
	}
}
