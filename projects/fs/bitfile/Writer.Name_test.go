package bitfile

import (
	"os"
	"testing"
)

/*
 * Writer.Name() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines the unit tests for the Writer.Name() method.
 */

func TestWriter_Name(t *testing.T) {
	var err error
	var bitfile Writer

	if bitfile.Name() != "" {
		t.Fatal("expect empty string when file not open")
	}
	bitfile.file, err = os.CreateTemp("", "Writer.Name.test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if bitfile.Name() == "" {
		t.Fatal("unexpected empty file name")
	}
	if bitfile.file.Name() != bitfile.Name() {
		t.Fatal("filename mismatch")
	}
}
