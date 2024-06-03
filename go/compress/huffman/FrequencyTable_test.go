package huffman

import "testing"

func TestFrequencyTableMap(t *testing.T) {
	var f FrequencyTable
	if f != nil {
		t.Fatal("Frequency Table should be nil")
	}
	f = make(FrequencyTable)
	if len(f) != 0 {
		t.Fatal("Frequency Table should be empty")
	}
	f[0x01] = 1
	if len(f) != 1 {
		t.Fatal("Frequency Table should have one entry")
	}
}
