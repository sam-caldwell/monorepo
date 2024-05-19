package directory

import (
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"testing"
)

func TestExistsp(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := "/tmp/test_dir"
	err := os.Mkdir(tempDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer func() { _ = os.RemoveAll(tempDir) }()

	// Test when directory exists
	exists := Existsp(&tempDir)
	if !exists {
		t.Fatalf("Expected directory to exist, but it does not.")
	}

	// Test when directory does not exist
	nonExistentDir := "/tmp/non_existent_dir"
	exists = Existsp(&nonExistentDir)
	if exists {
		t.Fatalf("Expected directory to not exist, but it does.")
	}
}
