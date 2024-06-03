package huffman

import "testing"

func TestCodeMap_Map(t *testing.T) {
	var c CodeMap
	if c != nil {
		t.Fatalf("expect uninitialized CodeMap to be nil")
	}
	c = make(CodeMap)
	if c == nil {
		t.Fatalf("expect non-nil CodeMap")
	}
	c[0xFF] = []byte("test")
	if string(c[0xFF]) != "test" {
		t.Fatalf("expect 'test' but got %v", string(c[0xFF]))
	}
}
