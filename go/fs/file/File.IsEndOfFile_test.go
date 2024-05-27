package file

import (
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
	"path/filepath"
	"testing"
)

func TestFile_IsEndOfFile(t *testing.T) {
	var (
		f                File
		name             string
		expectedFileName = filepath.Join("/tmp/", uuid.New().String())
		expectedValue    = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F}
		expectedSize     = int64(len(expectedValue))
		pos              int64
	)
	t.Cleanup(func() {
		_ = f.handle.Close()
		_ = os.Remove(name)
	})
	t.Run("test uninitialized file", func(t *testing.T) {
		if f.handle != nil {
			t.Fatal("handle is not nil")
		}
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
	t.Run("test that the file exists", func(t *testing.T) {
		if !Existsp(&name) {
			t.Fatal("file does not exist as expected")
		}
	})
	t.Run("Test write to the file", func(t *testing.T) {
		n, err := f.handle.Write(expectedValue)
		if err != nil {
			t.Fatal(err)
		}
		if int64(n) != expectedSize {
			t.Fatal("written length mismatch")
		}
	})
	t.Run("Test GetPosition()", func(t *testing.T) {
		var err error
		pos, err = f.GetPosition()
		if err != nil {
			t.Fatal(err)
		}
		if pos != expectedSize {
			t.Fatal("position mismatch")
		}
	})
	t.Run("Test IsEndOfFile()", func(t *testing.T) {
		eof, err := f.IsEndOfFile()
		if err != nil {
			t.Fatal(err)
		}
		if eof {
			t.Fatal("end of file mismatch")
		}
	})
}
