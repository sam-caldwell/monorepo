package simple

import (
	"testing"
)

func TestSet_List(t *testing.T) {
	var set Set[int]

	for i := 1; i <= 10; i++ {
		if err := set.Add(i); err != nil {
			t.Fatalf("error %d", i)
		}
	}
	if set.Count() < 10 {
		t.Fatal("Expected count is wrong (initialized)")
	}
	thisList := set.List()
	if len(thisList) < 10 {
		t.Fatal("Expected count is wrong (after list)")
	}
}
