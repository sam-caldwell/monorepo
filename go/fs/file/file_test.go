package file

import "testing"

func TestFile(t *testing.T) {
	const expectedFileName = "/tmp/myFile.txt"
	var f File
	if err := f.Set(expectedFileName); err != nil {
		t.Fatal(err)
	}
	if n := f.Get(); n != expectedFileName {
		t.Fatal("expected file name not found")
	}
}
