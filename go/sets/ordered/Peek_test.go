package ordered

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"testing"
)

func TestSet_Peek(t *testing.T) {
	var set Set[int]

	t.Run("Test with empty set", func(t *testing.T) {
		if _, err := set.Peek(0); err == nil {
			t.Fatalf("expected error not found")
		} else {
			if err.Error() != errors.IndexOutOfRange {
				t.Fatalf("expected %s, got %s", errors.IndexOutOfRange, err.Error())
			}
		}
	})
	t.Run("Setup the test data", func(t *testing.T) {
		_ = set.Add(2)
		_ = set.Add(3)
		_ = set.Add(4)
		_ = set.Add(5)
		_ = set.Add(6)
	})
	t.Run("Test set within bounds0", func(t *testing.T) {
		out, err := set.Peek(0)
		if err != nil {
			t.Fatalf("expected no error, got %s", err.Error())
		}
		if out != 2 {
			t.Fatalf("expected %d, got %d", 2, out)
		}
	})
	t.Run("Test set within bounds1", func(t *testing.T) {
		out, err := set.Peek(1)
		if err != nil {
			t.Fatalf("expected no error, got %s", err.Error())
		}
		if out != 3 {
			t.Fatalf("expected %d, got %d", 3, out)
		}
	})
	t.Run("Test set out of bounds", func(t *testing.T) {
		if _, err := set.Peek(5); err == nil {
			t.Fatalf("expected error not found")
		} else {
			if err.Error() != errors.IndexOutOfRange {
				t.Fatalf("expected %s, got %s", errors.IndexOutOfRange, err.Error())
			}
		}
	})
}
