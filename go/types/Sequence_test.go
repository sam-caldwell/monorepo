package types

import "testing"

func TestSequence(t *testing.T) {
	var seq Sequence
	if int(seq) != 0 {
		t.Fatal("expected 0")
	}
	if c := seq.Current(); c != 0 {
		t.Fatal("expected 0")
	}
	if int(seq.Get()) != 1 {
		t.Fatal("expected 1")
	}
	if c := seq.Current(); c != 1 {
		t.Fatal("expected 1")
	}
	if int(seq.Get()) != 2 {
		t.Fatal("expected 2")
	}
	if c := seq.Current(); c != 2 {
		t.Fatal("expected 2")
	}
}
