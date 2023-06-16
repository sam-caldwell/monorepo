package orderedset

import "testing"

func TestOrderedSet_Push(t *testing.T) {
	var set Set
	if err := set.Push(0); err != nil {
		t.Fatal(err)
	}
	if err := set.Push(1); err != nil {
		t.Fatal(err)
	}
	if err := set.Push(2); err != nil {
		t.Fatal(err)
	}
	if err := set.Push(3); err != nil {
		t.Fatal(err)
	}
	if set.Count() != 4 {
		t.Fatal("count mismatch")
	}
}
