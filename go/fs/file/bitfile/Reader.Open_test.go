package bitfile

/*
 * Reader.Open() tests
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

import (
	"os"
	"testing"
)

// TestBitFile_Open - create a test file, test the Reader.OpenRead() method and then clean up afterward.
func TestBitFile_Open(t *testing.T) {
	const testData = "this is a test"
	var testFile string
	testFileName := "/tmp/TestBitFile_Open.txt"

	t.Cleanup(func() {
		if err := os.Remove(testFileName); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create test file", func(t *testing.T) {
		var err error
		var f *os.File
		if f, err = os.Create(testFileName); err != nil {
			t.Fatal(err)
		}
		testFile = f.Name()
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
	t.Run("Perform the test of the bitfile.OpenRead() method", func(t *testing.T) {
		var f Reader
		if err := f.Open(&testFile, MinimumBlockSize); err != nil {
			t.Fatalf("failed to open file (%s): %v", testFile, err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				t.Fatal(err)
			}
		}()
	})
}
