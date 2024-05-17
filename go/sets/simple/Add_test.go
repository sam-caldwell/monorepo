package simple

import (
	"testing"
)

func TestSet_Add(t *testing.T) {
	t.Run("simple test with no initialization", func(t *testing.T) {
		var set Set[int]
		if set.data != nil {
			t.Fatal("expect nil data map at this point.")
		}
		if err := set.Add(1); err != nil {
			t.Fatalf("simple test: %v", err)
		}
		if err := set.Add(2); err != nil {
			t.Fatalf("simple test: %v", err)
		}
	})
	t.Run("simple test with prior initialization", func(t *testing.T) {
		var set Set[int]
		set.Init()
		if err := set.Add(1); err != nil {
			t.Fatalf("simple test: %v", err)
		}
		if err := set.Add(2); err != nil {
			t.Fatalf("simple test: %v", err)
		}
	})
	t.Run("test with different types", func(t *testing.T) {
		t.Run("test int", func(t *testing.T) {
			var set Set[int]
			if err := set.Add(1); err != nil {
				t.Fatalf("simple test: %v", err)
			}
		})
		t.Run("test uint", func(t *testing.T) {
			var set Set[uint]
			if err := set.Add(99); err != nil {
				t.Fatalf("simple test: %v", err)
			}
		})
		t.Run("test string", func(t *testing.T) {
			var set Set[string]
			if err := set.Add("data"); err != nil {
				t.Fatalf("simple test: %v", err)
			}
			if err := set.Add("data2"); err != nil {
				t.Fatalf("simple test: %v", err)
			}
		})
	})
}
