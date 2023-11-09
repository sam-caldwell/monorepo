package bitfile

/*
 * Reader.Read() test
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Tests for the bitfile Reader.Read() method
 */

import (
	"bytes"
	"os"
	"testing"
)

// TestReader_Read - create a test file and write some test data, then test the Reader.Read() method.
func TestReader_Read(t *testing.T) {
	testFileName := "/tmp/TestReader_Read.txt"
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

		var b Reader

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

	t.Run("Create an undersized block and read all testData", func(t *testing.T) {

		var b Reader

		if err := b.Open(&testFileName); err != nil {
			t.Fatal(err)
		}

		block, err := b.Read()

		if err != nil {
			t.Fatal(err)
		}

		if data := block.ReadBytes(uint(len(testData) - 1)); err != nil {
			t.Fatal(err)
		} else {
			if bytes.Equal(data, testData[0:uint(len(testData)-1)]) {
				t.Fatal("test data n-1 mismatch")
			}
		}

	})

}
