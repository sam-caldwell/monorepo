package file

import (
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"testing"
)

func TestFile_Delete(t *testing.T) {
	var (
		f                File
		expectedFileName = filepath.Join("/tmp/", uuid.New().String())
	)
	t.Run("Test uninitialized state", func(t *testing.T) {
		if f.handle != nil {
			t.Fatalf("file handle should be nil")
		}
	})
	t.Run("Create a test file", func(t *testing.T) {
		var err error
		f.handle, err = os.OpenFile(expectedFileName, os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			t.Fatal("error opening file", err)
		}
		if _, err = f.handle.Write([]byte("test")); err != nil {
			t.Fatal("error writing to file", err)
		}
		if err = f.handle.Close(); err != nil {
			t.Fatal("error closing file", err)
		}
	})
	t.Run("Delete a test file", func(t *testing.T) {
		if err := f.Delete(); err != nil {
			t.Fatal(err)
		}
	})
}
