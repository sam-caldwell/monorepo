package simple

import (
	"fmt"
	"testing"
)

func TestSet_ListString(t *testing.T) {
	var set Set
	_ = set.Add(0)
	_ = set.Add(1)
	_ = set.Add(2)
	_ = set.Add(3)
	_ = set.Add(4)
	_ = set.Add(5)
	for i, actual := range set.ListString(true) {
		if expected := fmt.Sprintf("%v", i); actual != expected {
			t.Fatalf("failed on %v actual:%v expected:%v", i, actual, expected)
		}
	}

}
