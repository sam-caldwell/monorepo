package ordered

import "testing"

func TestSet_Add(t *testing.T) {
	t.Run("integer test", func(t *testing.T) {
		set := Set[int]{data: []int{1, 2, 3}}

		t.Run("AddNonDuplicate", func(t *testing.T) {
			err := set.Add(4)
			if err != nil {
				t.Errorf("Expected nil error, got %v", err)
			}
		})

		t.Run("AddDuplicate", func(t *testing.T) {
			err := set.Add(3)
			expectedErr := "duplicate entry"
			if err == nil || err.Error() != expectedErr {
				t.Errorf("Expected error '%s', got %v", expectedErr, err)
			}
		})
	})
	t.Run("string test", func(t *testing.T) {
		var set Set[string]

		t.Run("AddFirstElement", func(t *testing.T) {
			err := set.Add("1")
			if err != nil {
				t.Errorf("Expected nil error, got %v", err)
			}
		})
		t.Run("AddNonDuplicate", func(t *testing.T) {
			err := set.Add("2")
			if err != nil {
				t.Errorf("Expected nil error, got %v", err)
			}
		})

		t.Run("AddDuplicate", func(t *testing.T) {
			err := set.Add("1")
			expectedErr := "duplicate entry"
			if err == nil || err.Error() != expectedErr {
				t.Errorf("Expected error '%s', got %v", expectedErr, err)
			}
		})
	})
}
