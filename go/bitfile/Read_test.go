package bitfile

/*
 * CRSCE bitfile
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
	const expected = `0111010001101000011010010111001100100000011010010111001100100000011000010010000001110100011001010111001101110100`
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

	//Perform the test of the bitfile.Open() method.
	func() {
		var f BitFile
		defer f.Close()

		//Open the test file.
		if err := f.Open(&testFile); err != nil {
			t.Fatalf("Open() failed: %v", err)
		}

		//Create a byte array we will use to store our data
		actual := ""

		sz := len([]byte(testData))
		//t.Logf("Size: %d", sz)
		//Read bit stream
		for i := 0; i < 8*sz; i++ {
			bit, err := f.Read()
			if err != nil {
				t.Fatalf("Error reading at %d", i)
			}
			if bit {
				actual += "1"
			} else {
				actual += "0"
			}
		}
		if actual != expected {
			t.Fatalf("actual read does not equal expected\n  actual: %s\nexpected: %s\n", actual, expected)
		}
	}()
}
