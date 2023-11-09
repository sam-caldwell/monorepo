package bitfile

/*
 * Bitfile Close() test
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Test the bitfile close method
 */

import (
	"os"
	"testing"
)

// TestBitFile_Close - create a test file, test the Reader.Close() method and then clean up afterward.
func TestBitFile_Close(t *testing.T) {

	var testFile string

	const testData = "this is a test"

	//Create test file
	func() {
		var err error
		var f *os.File
		if f, err = os.CreateTemp("", "TestBitFile_Close.*.txt"); err != nil {
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
		if err := os.Remove(testFile); err != nil {
			t.Fatalf("cleanup failed to delete test file: %v", err)
		}
	}()

	//Perform a .Close() with a nil file handle.  Expect no error
	func() {
		var f Reader
		if err := f.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	//Perform the test of the bitfile.Close() method.
	func() {
		var f Reader
		if err := f.Open(&testFile); err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := f.Close(); err != nil {
				t.Fatal(err)
			}
		}()
	}()

	//Perform the test of the bitfile.Close() method with double-close
	func() {
		var f Reader
		if err := f.Open(&testFile); err != nil {
			t.Fatal(err)
		}
		defer func() {
			var f Reader
			if err := f.Close(); err != nil {
				t.Fatal(err)
			}
		}()
	}()
}
