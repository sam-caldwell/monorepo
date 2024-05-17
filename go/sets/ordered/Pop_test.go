package ordered

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"testing"
)

func TestSet_Pop(t *testing.T) {
	var set Set[int]

	t.Run("pop from empty set", func(t *testing.T) {
		out, err := set.Pop()
		if out != 0 {
			t.Fatal("expected 0, got ", out)
		}
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if err.Error() != errors.EmptySet {
			t.Fatal("expected EmptySet, got ", err)
		}
	})
	t.Run("add data to set", func(t *testing.T) {
		_ = set.Add(0)
		_ = set.Add(1)
		_ = set.Add(2)
		_ = set.Add(3)
		_ = set.Add(4)
	})
	t.Run("pop from set", func(t *testing.T) {
		out, err := set.Pop()
		if err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
		if out != 0 {
			t.Fatalf("expected 0, got %d", out)
		}
	})
	t.Run("verify size after pop", func(t *testing.T) {
		if len(set.data) != 4 {
			t.Fatalf("expected 4, got %d", len(set.data))
		}
	})
	t.Run("pop from set", func(t *testing.T) {
		out, err := set.Pop()
		if err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
		if out != 1 {
			t.Fatalf("expected 1, got %d", out)
		}
	})
	t.Run("verify size after pop", func(t *testing.T) {
		if len(set.data) != 3 {
			t.Fatalf("expected 3, got %d", len(set.data))
		}
	})
	t.Run("pop from set", func(t *testing.T) {
		out, err := set.Pop()
		if err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
		if out != 2 {
			t.Fatalf("expected 2, got %d", out)
		}
	})
}
