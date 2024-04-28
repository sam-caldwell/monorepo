package file

import (
	"bytes"
	"os"
	"testing"
)

func TestFileCreate(t *testing.T) {
	fileName := "/tmp/testFile"
	const (
		expectedContent = "test content"
		expectedMode    = 0600
	)

	t.Cleanup(func() {
		if err := os.RemoveAll(fileName); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Create test file", func(t *testing.T) {
		if err := Create(fileName, []byte(expectedContent), expectedMode); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("confirm file exists", func(t *testing.T) {
		if !Exists(fileName) {
			t.Fatal("file not found")
		}
	})
	t.Run("confirm file contents", func(t *testing.T) {
		content, err := os.ReadFile(fileName)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(content, []byte(expectedContent)) {
			t.Fatal("content mismatch")
		}
		if info, err := os.Stat(fileName); err != nil {
			t.Fatal(err)
		} else {
			if info.Mode() != expectedMode {
				t.Fatal("mode mismatch")
			}
		}
	})
}
