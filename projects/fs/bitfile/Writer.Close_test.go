package bitfile

/*
 * Writer.Close() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Unit test for the Writer.Close() method
 */

import (
	"os"
	"testing"
)

func TestWriter_Close(t *testing.T) {
	const tempFileName = "/tmp/TestWriter_Close.txt"
	var w Writer
	var err error
	if w.file, err = os.Create(tempFileName); err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = w.file.Close()
		_ = os.Remove(tempFileName)
	}()

	if _, err = w.file.Write([]byte("test data")); err != nil {
		t.Fatal(err)
	}

	if err = w.Close(); err != nil {
		t.Fatal(err)
	}

}
