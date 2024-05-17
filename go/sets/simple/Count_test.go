package simple

import "testing"

func TestSet_Count(t *testing.T) {

	t.Run("Test Count method uninitialized", func(t *testing.T) {
		var set Set[int]
		// Verify that the initial count is 0
		if count := set.Count(); count != 0 {
			t.Errorf("Expected count 0, got: %d", count)
		}
	})
	t.Run("Test Count method initialized", func(t *testing.T) {
		var set Set[int]
		// Verify that the initial count is 0
		if count := set.Count(); count != 0 {
			t.Errorf("Expected count 0, got: %d", count)
		}

		// Add items to the set
		_ = set.Add(1)
		_ = set.Add(2)
		_ = set.Add(3)
		_ = set.Add(4)

		// Verify that the count is updated correctly
		if count := set.Count(); count != 4 {
			t.Errorf("Expected count 4, got: %d", count)
		}

		if err := set.Add(5); err != nil {
			t.Fatal(err)
		}

		// Verify that the count is updated correctly again
		if count := set.Count(); count != 5 {
			t.Errorf("Expected count 5, got: %d", count)
		}
	})

}
