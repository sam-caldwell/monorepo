package ordered

import "testing"

func TestSet_Empty(t *testing.T) {
	var set Set

	_ = set.Add(1)
	_ = set.Add(2)
	_ = set.Add(3)
	_ = set.Add(4)
}
