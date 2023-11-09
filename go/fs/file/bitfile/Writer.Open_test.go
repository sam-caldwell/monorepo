package bitfile

/*
 * Writer.Open() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Define unit test for the Writer.Open() method
 */

import (
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
	"testing"
)

func TestWriter_Open(t *testing.T) {
	var b Writer
	tempFileName := "/tmp/TestBitFile_Create.txt"
	if err := b.Open(&tempFileName); err != nil {
		t.Fatal(err)
	}
	if file.Exists(tempFileName) {
		if err := os.Remove(tempFileName); err != nil {
			t.Fatal(err)
		}
	}
	defer func() {
		if err := b.Close(); err != nil {
			t.Fatal(err)
		}
		if err := os.Remove(tempFileName); err != nil {
			t.Fatal(err)
		}
	}()
}
