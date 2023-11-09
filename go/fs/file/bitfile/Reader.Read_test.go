package bitfile

/*
 * bitfile.Read() test
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

import (
	"bytes"
	"os"
	"testing"
)

// TestBitFile_Read - create a test file and write some test data, then test the BitFile.Read() method.
func TestBitFile_Read(t *testing.T) {
	testFileName := "/tmp/TestBitFile_Read.txt"
	testData := []byte("this is a test")

	func() {
		//Create test file with data.
		if err := os.WriteFile(testFileName, testData, 0644); err != nil {
			t.Fatal(err)
		}
	}()
	defer func() {
		if err := os.Remove(testFileName); err != nil {
			t.Fatal(err)
		}
	}()

	t.Run("Create an oversize block and read all testData", func(t *testing.T) {
		//Create an oversize block
		var b BitFile
		if err := b.Open(&testFileName); err != nil {
			t.Fatal(err)
		}
		block, err := b.Read()
		if err != nil {
			t.Fatal(err)
		}
		if data := block.ReadBytes(0); !bytes.Equal(data, testData) {
			t.Fatal("test data mismatch")
		}
	})

}
