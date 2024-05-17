package simple

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"testing"
)

func TestSet_GetFirst(t *testing.T) {
	t.Run("GetFirst() test data with initialized set", func(t *testing.T) {
		var set Set[int]
		set.Init()
		if _, err := set.GetFirst(); err == nil {
			if err.Error() != errors.EmptySet {
				t.Fatal("error mismatch")
			} else {
				t.Errorf("Set.GetFirst() unexpected error: %v", err)
			}
		}
	})
	t.Run("GetFirst() test data with data", func(t *testing.T) {
		var set Set[int]
		set.Init()
		_ = set.Add(1)
		_ = set.Add(2)
		_ = set.Add(3)
		_ = set.Add(4)
		if v, err := set.GetFirst(); err != nil {
			t.Fatal(err)
		} else {
			if !set.Has(v) {
				t.Fatalf("GetFirst() should have returned a value in the set")
			}
		}
	})
}
