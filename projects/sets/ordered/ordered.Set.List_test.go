package ordered

import (
	"testing"
)

func TestOrderedSet_List(t *testing.T) {
	var set Set

	for i := 0; i < 10; i++ {
		if err := set.Add(i); err != nil {
			t.Fatalf("error %d", i)
		}
		if set.Count() != i+1 {
			t.Fatalf("Adding failed.  count does not track.")
		}
	}

	if set.Count() < 10 {
		t.Fatal("Expected count is wrong (initialized)")
	}

	thisList := set.List()

	if len(thisList) < 10 {
		t.Fatalf("Expected count is wrong (after list): (count:%d):%v", len(thisList), thisList)
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
