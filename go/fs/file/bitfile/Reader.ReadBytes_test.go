package bitfile

/*
 * Reader.ReadBytes() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Unit test for the Reader.ReadBytes() method
 */
import (
	"bytes"
	"os"
	"testing"
)

func TestReadBytes(t *testing.T) {
	tempFileName := "/tmp/TestReadBytes.txt"
	expected := []byte("This is a test message")

	func() {
		// Create a temporary file with some content
		if err := os.WriteFile(tempFileName, expected, 0644); err != nil {
			t.Fatal(err)
		}
	}()

	defer func() {
		if err := os.Remove(tempFileName); err != nil {
			t.Fatal(err)
		}
	}()

	// Happy path validation
	t.Run("Happy path: file reads okay", func(t *testing.T) {
		var b Reader
		if err := b.Open(&tempFileName); err != nil {
			t.Fatal(err)
		}

		n := uint(len(expected))
		blk, err := b.ReadBytes(n)
		if err != nil {
			t.Fatal(err)
		}
		if blk.Size() != n {
			t.Fatal("block size mismatch")
		}

		if actual := blk.ReadBytes(0); !bytes.Equal(actual, expected) {
			t.Fatal("data read mismatch(0)")
		}
		if actual := blk.ReadBytes(10); !bytes.Equal(actual, expected[0:10]) {
			t.Fatal("data read mismatch(10)")
		}
	})
	t.Run("Sad Path: file read fails when not open", func(t *testing.T) {
		var b Reader
		if err := b.Open(&tempFileName); err != nil {
			t.Fatal(err)
		}
		if err := b.Close(); err != nil {
			t.Fatal(err)
		}
		if _, err := b.ReadBytes(0); err == nil {
			t.Fatal("expected error")
		}
	})
}
