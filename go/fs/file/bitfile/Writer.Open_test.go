package bitfile

/*
 * Reader.Create() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Define unit test for the Create() method
 */

import (
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
	"testing"
)

func TestBitFile_Create(t *testing.T) {
	var b Reader
	tempFileName := "/tmp/TestBitFile_Create.txt"
	if err := b.Create(&tempFileName); err != nil {
		t.Fatal(err)
	}
	if !file.Exists(tempFileName) {
		t.Fatalf("failed to create file (%s)", tempFileName)
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
