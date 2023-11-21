package bitfile

/*
 * CRSCE Reader.WriteBytes() Test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Unit test for the Reader.WriteBytes() method which performs a bitwise buffered write to disk.
 */

import (
	"fmt"
	"testing"
)

func TestBitFile_WriteBytes(t *testing.T) {
	var target Writer

	t.Run("Setup the writer", func(t *testing.T) {
		target.buffer = make([]byte, 10)
		target.bufferPos = 0
		target.bitPos = 0
		target.file = nil //we do not want to open a file for this test.
	})
	t.Run("Write a set of 0xFF bytes", func(t *testing.T) {
		data := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
		if err := target.WriteBytes(data); err != nil {
			t.Fatalf("WriteBytes() failed. err: %v", err)
		}
	})

	t.Run("Verify all bytes", func(t *testing.T) {
		for i := uint16(0); i < uint16(len(target.buffer)); i++ {
			t.Run(fmt.Sprintf("Verify Byte %d", i), func(t *testing.T) {
				if actual := target.buffer[i]; actual != 0xFF {
					t.Fatalf("byte %d error (actual: %02x)", i, actual)
				}
			})
		}
	})

	t.Run("Reset the writer", func(t *testing.T) {
		target.buffer = make([]byte, 10)
		target.bufferPos = 0
		target.bitPos = 0
		target.file = nil //we do not want to open a file for this test.
	})

	t.Run("Verify all bytes", func(t *testing.T) {
		for i := uint16(0); i < uint16(len(target.buffer)); i++ {
			t.Run(fmt.Sprintf("Verify Byte %d", i), func(t *testing.T) {
				if actual := target.buffer[i]; actual != 0x00 {
					t.Fatalf("byte %d error (actual: %02x)", i, actual)
				}
			})
		}
	})

	t.Run("Write a set of sequential bytes", func(t *testing.T) {
		data := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09}
		if err := target.WriteBytes(data); err != nil {
			t.Fatalf("WriteBytes() failed. err: %v", err)
		}
	})

	t.Run("Verify sequential bytes", func(t *testing.T) {
		for i := uint16(0); i < uint16(len(target.buffer)); i++ {
			t.Run(fmt.Sprintf("Verify Byte %d", i), func(t *testing.T) {
				if actual := target.buffer[i]; actual != byte(i) {
					t.Fatalf("byte %d error (actual: %02x)", i, actual)
				}
			})
		}
	})
}
