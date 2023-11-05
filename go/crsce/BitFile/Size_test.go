package bitfile

/*
 * CRSCE BitFile size method test
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Test the BitFile.Size() method
 */

import (
	"os"
	"testing"
)

func TestBitFile_Size(t *testing.T) {
	const testData = "this is a test"
	var testFile string
	//Create test file
	func() {
		var err error
		var f *os.File

		if f, err = os.CreateTemp("", "TestBitFile_Read.*.txt"); err != nil {
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
	var f BitFile
	if err := f.Open(&testFile); err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	actual, err := f.Size()
	if err != nil {
		t.Fatalf("error retrieving size: %v", err)
	}
	if actual != int64(len(testData)) {
		t.Fatalf("size mismatch.  (actual: %d, expected: %d)", actual, len(testData))
	}
}
