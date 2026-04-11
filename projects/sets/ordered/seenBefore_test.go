package ordered

import "testing"

func TestSet_seenBefore(t *testing.T) {
	var set Set[int]

	t.Run("test on empty set", func(t *testing.T) {
		n := 0
		if ok := set.seenBefore(&n); ok {
			t.Fatal("expected false")
		}
	})
	t.Run("add data", func(t *testing.T) {
		_ = set.Add(0)
		_ = set.Add(1)
		_ = set.Add(2)
		_ = set.Add(3)
		_ = set.Add(4)
	})
	t.Run("test on existing set (found)", func(t *testing.T) {
		n := 0
		if ok := set.seenBefore(&n); !ok {
			t.Fatal("expected true")
		}
	})
	t.Run("test on existing set (not found)", func(t *testing.T) {
		n := 5
		if ok := set.seenBefore(&n); ok {
			t.Fatal("expected false")
		}
	})

}
