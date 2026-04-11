package ordered

import "testing"

func TestSet_Push(t *testing.T) {
	var set Set[int]
	t.Run("push to empty list", func(t *testing.T) {
		if err := set.Push(0); err != nil {
			t.Error(err)
		}
	})
	t.Run("push to non-empty list", func(t *testing.T) {
		if err := set.Push(1); err != nil {
			t.Error(err)
		}
		if err := set.Push(2); err != nil {
			t.Error(err)
		}
		if err := set.Push(3); err != nil {
			t.Error(err)
		}
		if err := set.Push(4); err != nil {
			t.Error(err)
		}
	})
	t.Run("test count", func(t *testing.T) {
		if len(set.data) != 5 {
			t.Error("wrong count")
		}
	})
}
