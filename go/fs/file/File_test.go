package file

import "testing"

func TestFile(t *testing.T) {
	const expectedFileName = "/tmp/myFile.txt"
	var f File
	if f.handle != nil {
		t.Fatal("handle expects to be nil")
	}
	f.lock.Lock()
	f.lock.Unlock()
}
