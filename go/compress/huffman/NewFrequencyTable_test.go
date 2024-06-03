package huffman

import "testing"

func TestNewFrequencyTable(t *testing.T) {
	table := NewFrequencyTable()
	if len(table) != 0 {
		t.Fatal("expect zero-length table")
	}
	table[0xFF] = 20
	if len(table) != 1 {
		t.Fatal("expect 1 table")
	}
	if table[0xFF] != 20 {
		t.Fatal("expect table 0xFF -> 20")
	}
}
