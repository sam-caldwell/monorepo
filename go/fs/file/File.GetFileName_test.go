package file

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
	"testing"
)

func TestFile_GetFileName(t *testing.T) {
	const expectedFileName = "/tmp/myFile.txt"
	var f File
	t.Run("Test with uninitialized file", func(t *testing.T) {
		if n := f.GetFileName(); n != words.EmptyString {
			t.Fatal("file name not to return if uninitialized")
		}
	})
	t.Run("initialize test", func(t *testing.T) {
		var err error
		if Exists(expectedFileName) {
			if err = os.Remove(expectedFileName); err != nil {
				t.Fatal("error removing file")
			}
		}
		f.handle, err = os.OpenFile(expectedFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			t.Fatal("failed to setup test")
		}
	})
	t.Cleanup(func() {
		_ = os.Remove(expectedFileName)
	})
	t.Run("test with initialized data", func(t *testing.T) {
		t.Run("Test with uninitialized file", func(t *testing.T) {
			if n := f.GetFileName(); n != expectedFileName {
				t.Fatal("expected file name not found")
			}
		})
	})
}
