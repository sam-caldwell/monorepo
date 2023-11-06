package bitfile

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestReadBytes(t *testing.T) {
	content := []byte("Hello, World!")

	// Create a temporary file with some content
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Error creating temporary file: %s", err)
	}
	defer func() {
		_ = tempFile.Close()
	}()

	if _, err := tempFile.Write(content); err != nil {
		t.Fatal(err)
	}

	// Create a BitFile using the temporary file
	bitFile := BitFile{
		file:   tempFile,
		buffer: make([]byte, len(content)),
	}
	//Move the file handle to the beginning of the file for our test to work
	_, _ = bitFile.file.Seek(0, 0)

	// Happy path validation
	t.Run("Happy path", func(t *testing.T) {
		n := len(content)
		blk, err := bitFile.ReadBytes(n)

		if err != nil {
			t.Errorf("Expected no error, got %v (content: %v)", err, blk.buffer)
		}

		if !bytes.Equal(blk.buffer, content) {
			t.Errorf("Expected buffer %v, got %v", content, blk.buffer)
		}
	})

	// Sad path validation (reading more bytes than available)
	t.Run("Sad path (reading more bytes than available)", func(t *testing.T) {
		n := len(content) + 1
		_, err := bitFile.ReadBytes(n)

		if err == nil || err.Error() != "EOF" {
			t.Errorf("Expected 'EOF' error, got %v", err)
		}
	})

	// Sad path validation (file not open)
	t.Run("Sad path (file not open)", func(t *testing.T) {
		if err := bitFile.file.Close(); err != nil {
			t.Fatal(err)
		}
		n := len(content)
		_, err := bitFile.ReadBytes(n)

		if err == nil || err.Error() != fmt.Sprintf("read %s: file already closed", tempFile.Name()) {
			t.Errorf("Expected 'file already closed' error, got %v", err)
		}
	})
}
