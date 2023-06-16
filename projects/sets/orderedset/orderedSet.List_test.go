package orderedset

import (
	"fmt"
	"testing"
)

func TestOrderedSet_List(t *testing.T) {
	var set Set

	for i := 0; i < 10; i++ {
		if err := set.Add(i); err != nil {
			t.Fatal(fmt.Sprintf("error %d", i))
		}
	}

	if set.Count() < 10 {
		t.Fatal("Expected count is wrong (initialized)")
	}

	thisList := set.List()

	if len(thisList) < 10 {
		t.Fatal("Expected count is wrong (after list)")
	}

	for i := 0; i < 10; i++ {
		item := set.Pop()
		if i != item {
			t.Fatalf("Mismatch at %d: %v", i, item)
		}
	}

	if len(thisList) == 0 {
		t.Fatal("Expected count to be 0")
	}
}
