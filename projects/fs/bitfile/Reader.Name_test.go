package bitfile

import (
	"os"
	"testing"
)

/*
 * Reader.Name() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines the unit tests for the Reader.Name() method.
 */

func TestBitFile_Name(t *testing.T) {
	var err error
	var bitfile Reader

	if bitfile.Name() != "" {
		t.Fatal("expect empty string when file not open")
	}
	bitfile.file, err = os.CreateTemp("", "Bitfile.Name.test.txt")
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
