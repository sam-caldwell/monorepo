package file

import (
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"testing"
)

func TestFile_Close(t *testing.T) {
	var f File
	t.Run("Test with uninitialized state", func(t *testing.T) {
		if f.handle != nil {
			t.Fatalf("file handle should be nil before test")
		}
		f.Close()
		if f.handle != nil {
			t.Fatalf("file handle should be nil after test")
		}
	})
	t.Run("Test with initialized state", func(t *testing.T) {
		var testFile = filepath.Join("/tmp/", uuid.New().String())
		t.Cleanup(func() {
			_ = os.Remove(testFile)
		})
		t.Run("Setup the test", func(t *testing.T) {
			var err error
			f.handle, err = os.OpenFile(testFile, os.O_CREATE|os.O_RDWR, 0600)
			if err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}
			if f.handle == nil {
				t.Fatalf("file handle should not be nil")
			}
		})
		t.Run("Run the test", func(t *testing.T) {
			f.Close()
		})
		t.Run("Evaluate the result", func(t *testing.T) {
			if f.handle != nil {
				t.Fatalf("file handle should be nil after .Close()")
			}
		})
	})
}
