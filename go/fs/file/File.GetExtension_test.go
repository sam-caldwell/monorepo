package file

import (
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFile_GetExtension(t *testing.T) {
	var (
		f                File
		name             string
		extension        = ".myExtension"
		expectedFileName = filepath.Join("/tmp/", uuid.New().String()) + extension
	)
	t.Run("test uninitialized state", func(t *testing.T) {
		if f.handle != nil {
			t.Fatal("file handle is nil")
		}
	})
	t.Run("test initialized state", func(t *testing.T) {
		t.Cleanup(func() {
			_ = f.handle.Close()
			_ = os.Remove(name)
		})
		t.Run("test create test file", func(t *testing.T) {
			var err error
			if err = f.Open(expectedFileName, os.O_CREATE|os.O_RDWR, 0600); err != nil {
				t.Fatal(err)
			}
			if f.handle == nil {
				t.Fatal("handle is nil")
			}
			name = f.handle.Name()
			if name == words.EmptyString {
				t.Fatal("name is empty")
			}
		})
		t.Run("test get the extension", func(t *testing.T) {
			if ext := f.GetExtension(); ext != strings.ToLower(extension) {
				t.Fatalf("extension is wrong.  Got: '%s', wanted: '%s'", ext, extension)
			}
		})
	})
}
