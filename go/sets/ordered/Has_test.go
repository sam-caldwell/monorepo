package ordered

import "testing"

func TestSet_Has(t *testing.T) {
	t.Run("Test Has on empty set", func(t *testing.T) {
		var set Set[int]
		if set.Has(1) {
			t.Fatal("Set should not have any elements so has should return false")
		}
	})
	t.Run("Test Has on non-empty set", func(t *testing.T) {
		var set Set[int]
		_ = set.Add(1)
		_ = set.Add(2)
		_ = set.Add(3)
		t.Run("item not found", func(t *testing.T) {
			if set.Has(4) {
				t.Fatal("Set should not have this element")
			}
		})
		t.Run("item found", func(t *testing.T) {
			if !set.Has(3) {
				t.Fatal("Set should have this element")
			}
		})
	})

}
