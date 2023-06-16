package simpleset

import (
	"fmt"
	"testing"
)

func TestSet_List(t *testing.T) {
	var set Set

	for i := 1; i <= 10; i++ {
		t.Log(i)
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
}
