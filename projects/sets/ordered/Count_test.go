package ordered

import "testing"

func TestSet_Count(t *testing.T) {
	t.Run("count empty set", func(t *testing.T) {
		var set Set[int]

		if count := set.Count(); count != 0 {
			t.Fatal("count failed 0.")
		}
	})
	t.Run("count 3-item set", func(t *testing.T) {
		set := Set[int]{data: []int{1, 2, 3}}

		if count := set.Count(); count != 3 {
			t.Fatal("count failed 3.")
		}
	})

}
