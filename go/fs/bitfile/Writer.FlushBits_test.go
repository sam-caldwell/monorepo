package bitfile

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestWriter_FlushBits(t *testing.T) {
	const bufferSize = 1024
	var target Writer
	testFileName := "/tmp/TestWriter_FlushBits.txt"
	t.Run("Writing single bit and flushing", func(t *testing.T) {
		t.Cleanup(func() {
			if target.file != nil {
				if err := target.file.Close(); err != nil {
					t.Fatalf("Failed to close: %v", err)
				}
				if err := os.Remove(testFileName); err != nil {
					t.Fatalf("failed to delete file: %v", err)
				}
				target.file = nil
			}
		})
		t.Run("Setup the writer", func(t *testing.T) {
			var err error
			target.buffer = make([]byte, bufferSize)
			target.bitPos = 0
			target.bufferPos = 0
			if target.file, err = os.Create(testFileName); err != nil {
				t.Fatalf("failed to create test file: %v", err)
			}
		})
		t.Run("write one bit", func(t *testing.T) {
			if err := target.WriteBit(0, 1); err != nil {
				t.Fatalf("Failed to write bit: %v", err)
			}
		})
		t.Run("pre-flush checks", func(t *testing.T) {
			t.Run("expect bit position to increment on WriteBit", func(t *testing.T) {
				if target.bitPos != 1 {
					// we expect bitPos to increment.
					t.Fatalf("pre-flush expects bitPos 1 after write but got %d", target.bitPos)
				}
			})
			t.Run("expect buffer position to be the same", func(t *testing.T) {
				if target.bufferPos != 0 {
					t.Fatalf("pre-flush expects bufferPos 0 but got %d", target.bitPos)
				}
			})
			t.Run("expect buffer to contain 0x01 after writeBit()", func(t *testing.T) {
				if target.buffer[target.bufferPos] != 0x01 {
					t.Fatalf("pre-flush value expects 0x01")
				}
			})
			t.Run("expect file position to be unchanged", func(t *testing.T) {
				var err error
				var currentPos int64
				if currentPos, err = target.file.Seek(0, io.SeekCurrent); err != nil {
					t.Fatalf("pre-flush error getting current file position: %v", err)
				}
				if currentPos != 0 {
					t.Fatalf("pre-flush error.  expected file position 0, got %d", currentPos)
				}
			})
		})
		t.Run("call FlushBits()", func(t *testing.T) {
			if err := target.FlushBits(); err != nil {
				t.Fatalf("failed to flush bits: %v", err)
			}
		})

		t.Run("post-flush checks", func(t *testing.T) {
			t.Run("expect bit position to reset on FlushBits()", func(t *testing.T) {
				if target.bitPos != 0 {
					// we expect bitPos to increment.
					t.Fatalf("post-flush expects bitPos 0 after write but got %d", target.bitPos)
				}
			})
			t.Run("expect buffer position to increment after FlushBits()", func(t *testing.T) {
				if target.bufferPos != 1 {
					t.Fatalf("post-flush expects bufferPos 1 but got %d", target.bitPos)
				}
			})
			t.Run("expect current buffer position to contain 0x00 after FlushBits()", func(t *testing.T) {
				if target.buffer[target.bufferPos] != 0x00 {
					t.Fatalf("post-flush value expects 0x00")
				}
			})
			t.Run("expect file position to be unchanged", func(t *testing.T) {
				var err error
				var currentPos int64
				if currentPos, err = target.file.Seek(0, io.SeekCurrent); err != nil {
					t.Fatalf("post-flush error getting current file position: %v", err)
				}
				if currentPos != 0 {
					t.Fatalf("post-flush error.  expected file position 0, got %d", currentPos)
				}
			})
		})
	})
	t.Run("Writing single byte will flush to disk one byte", func(t *testing.T) {
		t.Cleanup(func() {
			if target.file != nil {
				if err := target.file.Close(); err != nil {
					t.Fatalf("Failed to close: %v", err)
				}
				if err := os.Remove(testFileName); err != nil {
					t.Fatalf("failed to delete file: %v", err)
				}
				target.file = nil
			}
		})
		t.Run("Setup the writer", func(t *testing.T) {
			var err error
			target.buffer = make([]byte, bufferSize)
			target.bitPos = 0
			target.bufferPos = 0
			if target.file, err = os.Create(testFileName); err != nil {
				t.Fatalf("failed to create test file: %v", err)
			}
		})
		t.Run("write one byte to the buffer and setup the flush situation", func(t *testing.T) {
			target.buffer[target.bufferPos] = 0xFF
			expected := []byte{0xFF}
			if !bytes.Equal(target.buffer[0:1], expected) {
				t.Fatalf("buffer is not valid. %v", target.buffer)
			}
			target.bufferPos = uint16(len(target.buffer))
		})
		t.Run("call FlushBits()", func(t *testing.T) {
			if err := target.FlushBits(); err != nil {
				t.Fatalf("failed to flush bits: %v", err)
			}
		})
		t.Run("post-flush validation", func(t *testing.T) {
			t.Run("expect file position to advance", func(t *testing.T) {
				var err error
				var currentPos int64
				if currentPos, err = target.file.Seek(0, io.SeekCurrent); err != nil {
					t.Fatalf("error getting current file position: %v", err)
				}
				if currentPos != bufferSize {
					t.Fatalf("expected file position bufferSize, got %d", currentPos)
				}
			})
			t.Run("expect file size", func(t *testing.T) {
				stat, err := target.file.Stat()
				if err != nil {
					t.Fatalf("error getting file info.  %v", err)
				}
				if size := stat.Size(); size != bufferSize {
					t.Fatalf("size mismatch: %d", size)
				}
			})
			t.Run("expect file wrote 0xFF to the disk", func(t *testing.T) {
				if _, err := target.file.Seek(0, io.SeekStart); err != nil {
					t.Fatalf("Could not seek start of file. %v", err)
				}
				data := make([]byte, 1)
				if _, err := target.file.Read(data); err != nil {
					t.Fatalf("read failed. %v", err)
				}
				if !bytes.Equal(data, []byte{0xFF}) {
					t.Fatal("error on read-back")
				}
			})
		})
	})
}
