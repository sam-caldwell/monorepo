package file

import (
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"os"
	"testing"
)

func TestBinaryIo_Open(t *testing.T) {
	handle, err := file.CreateTempFile()
	defer func() {
		if handle == nil {
			return
		}
		_ = handle.Close()
		_ = os.Remove(handle.Name())
	}()
	if err != nil {
		t.Fatal(err)
	}
	if handle == nil {
		t.Fatal("handle should not be nil")
	}
	if err = handle.Sync(); err != nil {
		t.Fatalf("sync failed on file. %s", err.Error())
	}
	const message = "writeable file"
	if _, err = handle.Write([]byte(message)); err != nil {
		t.Fatalf("error writing to file. %s", err.Error())
	}
	if _, err = handle.Seek(0, 0); err != nil {
		t.Fatalf("error seeking file.  %s", err.Error())
	}
	buffer := make([]byte, len(message))
	if sz, err := handle.Read(buffer); err != nil {
		t.Fatalf("error reading from file. %s", err.Error())
	} else if sz != len(buffer) {
		t.Fatalf("error reading %d bytes", sz)
	}
}
