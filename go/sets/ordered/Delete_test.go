package ordered

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"testing"
)

func TestSet_Delete(t *testing.T) {
	t.Run("delete empty set", func(t *testing.T) {
		var set Set[int]
		if err := set.Delete(0); err == nil {
			t.Error("Expected error on delete0 of empty set")
		}
		if err := set.Delete(1); err == nil {
			t.Error("Expected error on delete1 of empty set")
		}
	})
	t.Run("delete from non-empty set", func(t *testing.T) {
		var set Set[int]
		t.Run("Setup test data", func(t *testing.T) {
			_ = set.Add(1)
			_ = set.Add(2)
			_ = set.Add(3)
		})
		t.Run("Delete item out of range", func(t *testing.T) {
			if err := set.Delete(3); err != nil {
				if err.Error() != errors.IndexOutOfRange {
					t.Fatalf("expected error mismatch")
				}
			} else {
				t.Fatal("expected error.  got none.")
			}
		})
		t.Run("Delete item.  expect no error", func(t *testing.T) {
			if err := set.Delete(0); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	})
}
