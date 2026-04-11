package bitfile

import (
	"testing"
)

/*
 * bitfile Writer.WriteBit() Test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Unit test for the Writer.WriteBit() method which performs a bitwise buffered write to disk.
 */

func TestBitFile_Write(t *testing.T) {
	var target Writer

	t.Run("Setup the writer", func(t *testing.T) {
		target.buffer = make([]byte, 10)
		target.bufferPos = 0
		target.bitPos = 0
		target.file = nil //we do not want to open a file for this test.
	})
	t.Run("Set bit 0", func(t *testing.T) {
		const a = 0b00000001
		const b = 0b00000001
		if err := target.WriteBit(0, a); err != nil {
			t.Fatalf("WriteBit() error %v", err)
		}
		if target.buffer[target.bufferPos] != b {
			t.Fatal("expected byte pattern not found")
		}
	})
	t.Run("Set bit 1", func(t *testing.T) {
		const a = 0b1111010
		const b = 0b00000011
		if err := target.WriteBit(1, a); err != nil {
			t.Fatalf("WriteBit() error %v", err)
		}
		if target.buffer[target.bufferPos] != b {
			t.Fatal("expected byte pattern not found")
		}
	})
	t.Run("Set bit 2", func(t *testing.T) {
		const a = 0b00100100
		const b = 0b00000111
		if err := target.WriteBit(2, a); err != nil {
			t.Fatalf("WriteBit() error %v", err)
		}
		if target.buffer[target.bufferPos] != b {
			t.Fatal("expected byte pattern not found")
		}
	})
	t.Run("Set bit 3", func(t *testing.T) {
		const a = 0b00001000
		const b = 0b00001111
		if err := target.WriteBit(3, a); err != nil {
			t.Fatalf("WriteBit() error %v", err)
		}
		if target.buffer[target.bufferPos] != b {
			t.Fatal("expected byte pattern not found")
		}
	})
	t.Run("Set bit 4", func(t *testing.T) {
		const a = 0b00010000
		const b = 0b00011111
		if err := target.WriteBit(4, a); err != nil {
			t.Fatalf("WriteBit() error %v", err)
		}
		if target.buffer[target.bufferPos] != b {
			t.Fatal("expected byte pattern not found")
		}
	})
	/*
	 * Clearing bits
	 */
	t.Run("Clear bit 0", func(t *testing.T) {
		const a = 0b00000000
		const b = 0b00011110
		if err := target.WriteBit(0, a); err != nil {
			t.Fatalf("WriteBit() error %v", err)
		}
		if actual := target.buffer[target.bufferPos]; actual != b {
			t.Fatalf("expected byte pattern not found\n"+
				"actual:   %08b\n"+
				"expected: %08b",
				actual, b)
		}
		if target.bitPos != 1 {
			t.Fatalf("bitPos mismatch at %d", target.bitPos)
		}
	})
	t.Run("Clear bit 1", func(t *testing.T) {
		const a = 0b00000000
		const b = 0b00011100
		if err := target.WriteBit(1, a); err != nil {
			t.Fatalf("WriteBit() error %v", err)
		}
		if actual := target.buffer[target.bufferPos]; actual != b {
			t.Fatalf("expected byte pattern not found\n"+
				"actual:   %08b\n"+
				"expected: %08b",
				actual, b)
		}
		if target.bitPos != 2 {
			t.Fatalf("bitPos mismatch at %d", target.bitPos)
		}
	})
	t.Run("Clear bit 2", func(t *testing.T) {
		const a = 0b00000000
		const b = 0b00011000
		if err := target.WriteBit(2, a); err != nil {
			t.Fatalf("WriteBit() error %v", err)
		}
		if actual := target.buffer[target.bufferPos]; actual != b {
			t.Fatalf("expected byte pattern not found\n"+
				"actual:   %08b\n"+
				"expected: %08b",
				actual, b)
		}
		if target.bitPos != 3 {
			t.Fatalf("bitPos mismatch at %d", target.bitPos)
		}
	})
	t.Run("Clear bit 3", func(t *testing.T) {
		const a = 0b00000000
		const b = 0b00010000
		if err := target.WriteBit(3, a); err != nil {
			t.Fatalf("WriteBit() error %v", err)
		}
		if actual := target.buffer[target.bufferPos]; actual != b {
			t.Fatalf("expected byte pattern not found\n"+
				"actual:   %08b\n"+
				"expected: %08b",
				actual, b)
		}
		if target.bitPos != 4 {
			t.Fatalf("bitPos mismatch at %d", target.bitPos)
		}
	})
	t.Run("Clear bit 4", func(t *testing.T) {
		const a = 0b00000000
		const b = 0b00000000
		if err := target.WriteBit(4, a); err != nil {
			t.Fatalf("WriteBit() error %v", err)
		}
		if actual := target.buffer[target.bufferPos]; actual != b {
			t.Fatalf("expected byte pattern not found\n"+
				"actual:   %08b\n"+
				"expected: %08b",
				actual, b)
		}
		if target.bitPos != 5 {
			t.Fatalf("bitPos mismatch at %d", target.bitPos)
		}
	})
	t.Run("set all bits", func(t *testing.T) {
		target.bitPos = 0
		const a = 0b11111111
		var b byte
		for i := byte(0); i < 8; i++ {
			b = b | (0b00000001 << i)
			if target.bitPos != i {
				t.Fatalf("bitPos mismatch at %d", target.bitPos)
			}
			if err := target.WriteBit(i, a); err != nil {
				t.Fatalf("WriteBit() error %v", err)
			}
			if actual := target.buffer[target.bufferPos]; actual != b {
				t.Fatalf("expected byte pattern not found\n"+
					"actual:   %08b\n"+
					"expected: %08b",
					actual, b)
			}
		}
	})
}
