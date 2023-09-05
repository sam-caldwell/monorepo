package directory

import (
	os2 "github.com/sam-caldwell/monorepo/go/projects/wrappers/os"
	"testing"
)

func TestExists(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := "/tmp/test_dir"
	err := os2.Mkdir(tempDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer func() { _ = os2.RemoveAll(tempDir) }()

	// Test when directory exists
	exists := Exists(tempDir)
	if !exists {
		t.Errorf("Expected directory to exist, but it does not.")
	}

	// Test when directory does not exist
	nonExistentDir := "/tmp/non_existent_dir"
	exists = Exists(nonExistentDir)
	if exists {
		t.Errorf("Expected directory to not exist, but it does.")
	}
}