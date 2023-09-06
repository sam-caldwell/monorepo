package simple

import "testing"

func TestSet_Empty(t *testing.T) {
	var set Set

	if set.data != nil {
		t.Fatal("Expected set.data to be nil initially")
	}

	if !set.Empty() {
		t.Fatal("Expected empty set initially")
	}
	set.Init()

	if !set.Empty() {
		t.Fatal("Expected initialized set to be empty")
	}
	if err := set.Add(1); err != nil {
		t.Fatal(err)
	}
	if set.Empty() {
		t.Fatal("Expected set to be non-empty after .Add()")
	}
}
