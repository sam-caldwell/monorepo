package file

import (
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"os"
	"testing"
)

func TestBinaryIo_Flush(t *testing.T) {
	handle, err := file.CreateTempFile()
	fileName := handle.Name()
	_ = handle.Close()

	defer func() {
		_ = os.Remove(fileName)
	}()

	var io BinaryIo
	if err = io.Open(fileName); err != nil {
		t.Fatalf("Fail to open file: %s", err.Error())
	}
	defer io.Close()
	io.Flush()
}
