package ordered

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"testing"
)

func TestSet_Insert(t *testing.T) {
	var set Set[int]
	t.Run("Insert on empty set (out of bounds)", func(t *testing.T) {
		if err := set.Insert(1, 10); err == nil {
			t.Fatalf("Insert on empty set (out of bounds) @ 10")
		} else {
			if err.Error() != errors.IndexOutOfRange {
				t.Fatalf("error mismatch")
			}
		}
	})
	t.Run("Insert on empty set (happy path)", func(t *testing.T) {
		if err := set.Insert(1, 0); err != nil {
			t.Fatalf("Insert on empty set (happy path) @ 0 (value 1): err:%v", err)
		}
	})
	t.Run("Add data to set", func(t *testing.T) {
		_ = set.Add(2) // pos: 1
		_ = set.Add(3) // pos: 2
		_ = set.Add(4) // pos: 3
		_ = set.Add(5) // pos: 4
		_ = set.Add(6) // pos: 5
	})
	t.Run("Insert on non-empty set (out of bounds, non-duplicate)", func(t *testing.T) {
		if err := set.Insert(10, 10); err == nil {
			t.Fatalf("Insert on empty set (out of bounds) @ 10")
		} else {
			if err.Error() != errors.IndexOutOfRange {
				t.Fatalf("error mismatch")
			}
		}
	})
	t.Run("Insert on non-empty set (out of bounds, duplicate)", func(t *testing.T) {
		if err := set.Insert(2, 10); err == nil {
			t.Fatalf("Insert on empty set (out of bounds) @ 10")
		} else {
			if err.Error() != errors.DuplicateEntry {
				t.Fatalf("error mismatch")
			}
		}
	})
	t.Run("Insert on non-empty set (happy path)", func(t *testing.T) {
		if err := set.Insert(99, 3); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("Verify set contents", func(t *testing.T) {
		expected := []int{1, 2, 3, 99, 4, 5, 6}
		for i, item := range expected {
			if item != set.data[i] {
				t.Fatalf("content mismatch at %d got %d but expected %d", i, set.data[i], item)
			}
		}
	})
}
