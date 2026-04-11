package bitfile

import "testing"

func TestWriter_setBit(t *testing.T) {
	var target Writer
	t.Run("Setup the writer", func(t *testing.T) {
		target.buffer = make([]byte, 10)
		target.bufferPos = 0
		target.bitPos = 0
		target.file = nil //we do not want to open a file for this test.
	})
	/*
	 * bit 0
	 */
	t.Run("Set bit 0", func(t *testing.T) {
		target.bitPos = 0
		target.setBit()
	})
	t.Run("Verify bit 0 is set", func(t *testing.T) {
		target.bitPos = 0
		mask := byte(1 << target.bitPos)
		if target.buffer[target.bufferPos]&mask == 0 {
			t.Fatal("expected set bit but it was clear")
		}
	})
	/*
	 * bit 1
	 */
	t.Run("Set bit 1", func(t *testing.T) {
		target.bitPos = 1
		target.setBit()
	})
	t.Run("Verify bit 0 is set", func(t *testing.T) {
		target.bitPos = 1
		mask := byte(1 << target.bitPos)
		if target.buffer[target.bufferPos]&mask == 0 {
			t.Fatal("expected set bit but it was clear")
		}
	})
	/*
	 * reset buffer to zero values
	 */
	t.Run("Verify all ***bytes*** are clear", func(t *testing.T) {
		for i := 0; i < len(target.buffer); i++ {
			target.buffer[i] = 0x00
		}
	})
	t.Run("set all bits", func(t *testing.T) {
		for i := byte(0); i < 8; i++ {
			target.bitPos = i
			target.setBit()
		}
	})
	t.Run("verify all bits are set", func(t *testing.T) {
		for i := byte(0); i < 8; i++ {
			var readMask = byte(1 << i)
			if bit := target.buffer[target.bufferPos]&readMask == 0; bit {
				t.Fatalf("bit %d is clear but expected set", i)
			}
		}
	})
	/*
	 * Walking set bits
	 */
	t.Run("Clear the buffer", func(t *testing.T) {
		for target.bufferPos = 0; target.bufferPos < uint16(len(target.buffer)); target.bufferPos++ {
			target.buffer[target.bufferPos] = 0x00
		}
	})
	t.Run("interleave set bits in the current byte across the buffer", func(t *testing.T) {
		for target.bufferPos = 0; target.bufferPos < uint16(len(target.buffer)); target.bufferPos += 1 {
			for target.bitPos = byte(0); target.bitPos < 8; target.bitPos += 2 {
				target.setBit()
			}
		}
	})
	t.Run("verify interleaved byte values", func(t *testing.T) {
		for target.bufferPos = 0; target.bufferPos < uint16(len(target.buffer)); target.bufferPos += 1 {
			if target.buffer[target.bufferPos] != 0b01010101 {
				t.Fatalf("interleaved bits mismatch at %d (value %08b)",
					target.bufferPos, target.buffer[target.bufferPos])
			}
		}
	})
}
