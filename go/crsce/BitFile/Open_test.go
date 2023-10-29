package bitfile

/*
 * CRSCE BitFile
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

import (
	"os"
	"testing"
)

// TestBitFile_Open - create a test file, test the BitFile.Open() method and then clean up afterward.
func TestBitFile_Open(t *testing.T) {
	var testFile string
	const testData = "this is a test"
	//Create test file
	func() {
		var err error
		var f *os.File
		if f, err = os.CreateTemp("", "TestBitFile_Open.*.txt"); err != nil {
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
	}()

	//Delete the test file when the test is done.
	defer func() {
		//t.Logf("deleting file")
		if err := os.Remove(testFile); err != nil {
			t.Fatalf("cleanup failed to delete test file: %v", err)
		}
		//t.Logf("deleted file: %s", testFile)
	}()

	//Perform the test of the BitFile.Open() method.
	func() {
		var f BitFile
		if err := f.Open(&testFile); err != nil {
			t.Fatalf("failed to open file (%s): %v", testFile, err)
		}
		defer f.Close()
		if f.buffer == nil {
			t.Fatal("expect buffer to be initialized")
		}
	}()
}
