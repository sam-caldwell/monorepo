package bitfile

/*
 * Reader.Read() test
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Tests for the bitfile Reader.Read() method
 */

import (
	"bytes"
	"github.com/sam-caldwell/monorepo/go/types/bitBlock"
	"os"
	"testing"
)

// TestReader_Read - create a test file and write some test data, then test the Reader.Read() method.
func TestReader_Read(t *testing.T) {
	const testDataSize = 8192
	testFileName := "/tmp/TestReader_Read.txt"

	// Create a sequence of byte values (ASCII) as test data
	testData := make([]byte, testDataSize)
	for i := 0; i < testDataSize; i++ {
		const numberLowerCase = 26
		const numberUpperCase = 26
		const numberDigits = 10
		const numberCharactersAndNumbers = numberLowerCase + numberUpperCase + numberDigits
		const startingAscii = 48
		testData[i] = byte(i % (numberCharactersAndNumbers + startingAscii))
	}

	t.Cleanup(func() {
		if err := os.Remove(testFileName); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Create test file with data", func(t *testing.T) {
		if err := os.WriteFile(testFileName, testData, 0644); err != nil {
			t.Fatal(err)
		}
		data, err := os.ReadFile(testFileName)
		if err != nil {
			t.Fatalf("Content verification failed: %v", err)
		}
		if !bytes.Equal(data, testData) {
			t.Fatalf("Content verification failed (mismatch)\n"+
				"actual:   %v\n"+
				"expected: %v", data, testData)
		}
	})

	t.Run("Read all testData into a block and compare", func(t *testing.T) {

		var b Reader
		var err error
		var block *bitBlock.Block

		if err = b.Open(&testFileName, testDataSize); err != nil {
			t.Fatalf("failed to open file: %v", err)
		}

		if block != nil {
			t.Fatal("expect nil block at this point")
		}

		if block, err = b.Read(); err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}

		if size := block.Size(); size != testDataSize {
			t.Fatalf("Block size mismatch")
		}

		expected := testData[:testDataSize]
		if actual := block.ReadBytes(0); !bytes.Equal(actual, expected) {
			t.Fatalf("test data mismatch\n"+
				"actual:  %v\n"+
				"expected %v",
				actual,
				expected)
		}
	})
	t.Run("Create a block of min size and read from the file only one block of data", func(t *testing.T) {

		var b Reader
		var err error
		var block *bitBlock.Block

		if err = b.Open(&testFileName, MinimumBlockSize); err != nil {
			t.Fatalf("failed to open file: %v", err)
		}

		if block != nil {
			t.Fatal("expect nil block at this point")
		}

		if block, err = b.Read(); err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}

		if size := block.Size(); size != MinimumBlockSize {
			t.Fatalf("Block size mismatch")
		}

		expected := testData[:MinimumBlockSize]
		if actual := block.ReadBytes(0); !bytes.Equal(actual, expected) {
			t.Fatalf("test data mismatch\n"+
				"actual:  %v\n"+
				"expected %v",
				actual,
				expected)
		}
	})
	t.Run("Create a block of min size and read from the file only one block at a time", func(t *testing.T) {

		var b Reader
		var err error
		var block *bitBlock.Block

		if err = b.Open(&testFileName, MinimumBlockSize); err != nil {
			t.Fatalf("failed to open file: %v", err)
		}

		if block != nil {
			t.Fatal("expect nil block at this point")
		}

		if block, err = b.Read(); err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}

		if size := block.Size(); size != MinimumBlockSize {
			t.Fatalf("Block size mismatch")
		}

		expected := testData[:MinimumBlockSize]
		if actual := block.ReadBytes(0); !bytes.Equal(actual, expected) {
			t.Fatalf("test data mismatch\n"+
				"actual:  %v\n"+
				"expected %v",
				actual,
				expected)
		}
		expected = testData[:MinimumBlockSize]
		if actual := block.ReadBytes(0); !bytes.Equal(actual, expected) {
			t.Fatalf("test data mismatch\n"+
				"actual:  %v\n"+
				"expected %v",
				actual,
				expected)
		}
		expected = testData[:MinimumBlockSize]
		if actual := block.ReadBytes(0); !bytes.Equal(actual, expected) {
			t.Fatalf("test data mismatch\n"+
				"actual:  %v\n"+
				"expected %v",
				actual,
				expected)
		}
	})
}
