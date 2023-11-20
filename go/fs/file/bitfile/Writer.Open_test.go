package bitfile

/*
 * Writer.Open() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Define unit test for the Writer.Open() method
 */

import (
	"os"
	"testing"
)

func TestWriter_Open(t *testing.T) {

	var b Writer
	testFileName := "/tmp/TestWriter_Open.txt"

	t.Cleanup(func() {
		if b.file == nil {
			return
		}
		if err := b.file.Close(); err != nil {
			t.Fatalf("close failed: %v", err)
		}
		if err := os.Remove(testFileName); err != nil {
			t.Fatalf("file remove failed: %v", err)
		}
	})

	t.Run("Open the test file (after making sure it doesn't exist)", func(t *testing.T) {
		_ = os.Remove(testFileName)
		if err := b.Open(&testFileName, 0); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Create test file and close it.  Then Open with the Writer.", func(t *testing.T) {
		var err error
		var f *os.File
		if f, err = os.Create(testFileName); err != nil {
			t.Fatal(err)
		}
		if err = f.Close(); err != nil {
			t.Fatal(err)
		}
		if err := b.Open(&testFileName, 0); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Create test file, write data, close it. Then open and verify truncation", func(t *testing.T) {
		var err error
		var f *os.File
		_ = os.Remove(testFileName)
		if f, err = os.Create(testFileName); err != nil {
			t.Fatal(err)
		}
		if err = f.Close(); err != nil {
			t.Fatal(err)
		}
		if err := b.Open(&testFileName, 0); err != nil {
			t.Fatal(err)
		}
		var n int
		var data []byte
		if _, err = b.file.Seek(0, 0); err != nil {
			t.Fatalf("failed to seec position 0,0: %v", err)
		}
		if n, err = b.file.Read(data); err != nil {
			t.Fatalf("failed to read. %v", err)
		}
		if len(data) != 0 || n != 0 {
			t.Fatal("expected 0 bytes read")
		}

	})
}
