package ordered

import "testing"

func TestSet_Empty(t *testing.T) {
	t.Run("count empty set", func(t *testing.T) {
		var set Set[int]

		if e := set.Empty(); !e {
			t.Fatal("not empty")
		}
	})
	t.Run("count non-empty set", func(t *testing.T) {
		set := Set[int]{data: []int{1, 2, 3}}

		if e := set.Empty(); e {
			t.Fatalf("empty count:%d", set.Count())
		}
	})

}
