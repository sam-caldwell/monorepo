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

// TestBitFile_Read - create a test file and write some test data, then test the BitFile.Read() method.
func TestBitFile_Read(t *testing.T) {
	var testFile string

	const testData = "this is a test"

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

	//Delete the test file when the test is done.
	defer func() {
		if err := os.Remove(testFile); err != nil {
			t.Fatalf("cleanup failed to delete test file: %v", err)
		}
	}()

	//Perform the test of the BitFile.Open() method.
	func() {
		var f BitFile
		defer f.Close()

		//Open the test file.
		if err := f.Open(&testFile); err != nil {
			t.Fatal(err)
		}

		if len([]byte(testData)) == 0 {
			t.Fatal("testData cannot be zero-length")
		}

		//Create a byte array we will use to store our data
		actual := make([]byte, len([]byte(testData)))

		if len(actual) == 0 {
			t.Fatal("actual must be more than 0 bytes")
		}

		//Read
		for i := 0; i < len([]byte(testData)); i++ {
			bit, err := f.Read()
			if err != nil {
				t.Fatalf("Error reading at %d", i)
			}
			if bit {
				actual[i] = 1
			} else {
				actual[i] = 0
			}
		}
	}()
}
