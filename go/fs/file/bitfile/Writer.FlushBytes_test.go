package bitfile

import "testing"

func TestWriter_FlushBuffer(t *testing.T) {
	var target Writer

	t.Run("Setup the writer", func(t *testing.T) {
		target.buffer = make([]byte, 10)
		target.bufferPos = 0
		target.bitPos = 0
		target.file = nil //we do not want to open a file for this test.
	})
	t.Run("Write a single byte", func(t *testing.T) {
		data := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
		if err := target.WriteBytes(data); err != nil {
			t.Fatalf("WriteBytes() failed. err: %v", err)
		}
	})
	//t.Run("Verify Bytes", func(t *testing.T) {
	//	if target.buffer[0] != 0xFF {
	//		t.Fatal("byte 0 error")
	//	}
	//})
}
