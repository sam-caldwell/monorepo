package ordered

import "testing"

func TestSet_Add(t *testing.T) {
	t.Run("integer test", func(t *testing.T) {
		var set Set[int]

		t.Run("AddNonDuplicate", func(t *testing.T) {
			err := set.Add(1)
			if err != nil {
				t.Fatalf("Expected nil error, got %v", err)
			}
			_ = set.Add(2)
			_ = set.Add(3)
			_ = set.Add(4)
			_ = set.Add(5)
			_ = set.Add(6)
		})

		t.Run("AddDuplicate", func(t *testing.T) {
			err := set.Add(3)
			expectedErr := "duplicate entry"
			if err == nil || err.Error() != expectedErr {
				t.Fatalf("Expected error '%s', got %v", expectedErr, err)
			}
		})
	})
	t.Run("string test", func(t *testing.T) {
		var set Set[string]
		t.Run("AddFirstElement", func(t *testing.T) {
			err := set.Add("1")
			if err != nil {
				t.Fatalf("Expected nil error, got %v", err)
			}
		})
		t.Run("AddNonDuplicate", func(t *testing.T) {
			err := set.Add("2")
			if err != nil {
				t.Fatalf("Expected nil error, got %v", err)
			}
		})

		t.Run("AddDuplicate", func(t *testing.T) {
			err := set.Add("1")
			expectedErr := "duplicate entry"
			if err == nil || err.Error() != expectedErr {
				t.Fatalf("Expected error '%s', got %v", expectedErr, err)
			}
		})
	})
}
