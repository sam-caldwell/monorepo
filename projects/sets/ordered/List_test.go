package ordered

import (
	"testing"
)

func TestSet_List(t *testing.T) {
	var set Set[int]
	t.Run("list empty set", func(t *testing.T) {
		result := set.List()
		if len(result) != 0 {
			t.Fatalf("expected empty list: %v", result)
		}
	})
	t.Run("add data to set", func(t *testing.T) {
		_ = set.Add(0)
		_ = set.Add(1)
		_ = set.Add(2)
		_ = set.Add(3)
		_ = set.Add(4)
	})
	t.Run("list non-empty set (check size)", func(t *testing.T) {
		result := set.List()
		if len(result) != len(set.data) {
			t.Fatalf("expected 5-item list: %v", result)
		}
	})
	t.Run("list non-empty set (check content)", func(t *testing.T) {
		result := set.List()
		for i, item := range result {
			if item != set.data[i] {
				t.Fatalf("data mismatch at %d", i)
			}
		}
	})
}
