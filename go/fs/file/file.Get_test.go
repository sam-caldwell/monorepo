package file

import "testing"

func TestFileGet(t *testing.T) {
	const expectedFileName = "myFile.txt"
	f := File(expectedFileName)
	if f.Get() != expectedFileName {
		t.Fatalf("get should return contents of the internal state")
	}

}
