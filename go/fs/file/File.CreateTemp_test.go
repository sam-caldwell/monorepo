package file

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
	"testing"
)

func TestFile_CreateTemp(t *testing.T) {
	var f File
	t.Run("test uninitialized file", func(t *testing.T) {
		if f.handle != nil {
			t.Fatal("handle is not nil")
		}
	})
	t.Run("test create temp file", func(t *testing.T) {
		var err error
		if err = f.CreateTemp(0600); err != nil {
			t.Fatal(err)
		}
		if f.handle == nil {
			t.Fatal("handle is nil")
		}
	})
	t.Run("test that the file exists", func(t *testing.T) {
		name := f.handle.Name()
		t.Cleanup(func() {
			_ = f.handle.Close()
			_ = os.Remove(name)
		})
		if name == words.EmptyString {
			t.Fatal("name is empty")
		}
		if !Exists(name) {
			t.Fatal("file does not exist as expected")
		}
	})
}
