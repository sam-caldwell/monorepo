package bitfile

/*
 * bitfile Reader.Size() tests
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Test the Reader.Size() method
 */

import (
	"os"
	"testing"
)

func TestBitFileReader_Size(t *testing.T) {

	const testData = "this is a test"
	testFileName := "/tmp/TestBitFileReader_Size.txt"

	t.Cleanup(func() {
		if err := os.Remove(testFileName); err != nil {
			t.Fatalf("file remove failed: %v", err)
		}
	})

	t.Run("Create test file", func(t *testing.T) {
		var err error
		var f *os.File

		if f, err = os.Create(testFileName); err != nil {
			t.Fatal(err)
		}
		n, err := f.Write([]byte(testData))
		if err != nil {
			t.Fatal("error writing test data")
		}
		if n != len([]byte(testData)) {
			t.Fatal("test data length mismatch.")
		}
		if err = f.Close(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Open file and read content", func(t *testing.T) {
		var f Reader

		if err := f.Open(&testFileName, MinimumBlockSize); err != nil {
			t.Fatalf("failed to open file: %v", err)
		}

		actual, err := f.Size()

		if err != nil {
			t.Fatalf("error retrieving size: %v", err)
		}

		if actual != uint64(len(testData)) {
			t.Fatalf("size mismatch.  (actual: %d, expected: %d)", actual, len(testData))
		}
	})

}
