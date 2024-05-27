package file

import (
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"testing"
)

func TestFile_struct(t *testing.T) {
	var expectedFileName = filepath.Join("/tmp/", uuid.New().String())
	var f File
	t.Run("pre-test", func(t *testing.T) {
		if f.handle != nil {
			t.Fatal("handle expects to be nil")
		}
	})
	t.Run("setup test", func(t *testing.T) {
		var err error
		f.handle, err = os.OpenFile(expectedFileName, os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			t.Fatal("error opening file", err)
		}
		if err = f.handle.Close(); err != nil {
			t.Fatal("error closing file", err)
		}
	})
	t.Run("test locks", func(t *testing.T) {
		f.lock.Lock()
		f.lock.Unlock()
	})
}
