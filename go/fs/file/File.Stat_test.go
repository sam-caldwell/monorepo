package file

import (
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestFile_Stat(t *testing.T) {
	var (
		f                File
		name             string
		expectedFileName = filepath.Join("/tmp/", uuid.New().String())
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
	t.Run("tests file stat", func(t *testing.T) {
		expected, err := os.Stat(name)
		if err != nil {
			t.Fatal(err)
		}
		actual, err := f.Stat()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Fatal("actual does not match expected file")
		}
	})
}
