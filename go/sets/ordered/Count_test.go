package ordered

import "testing"

func TestOrderedSet_Count(t *testing.T) {
	// Create a new set
	set := Set{}

	// Verify that the initial count is 0
	if count := set.Count(); count != 0 {
		t.Errorf("Expected count 0, got: %d", count)
	}

	// Add items to the set
	set.insert(1)
	set.insert(2)
	set.insert(3)

	// Verify that the count is updated correctly
	if count := set.Count(); count != 3 {
		t.Errorf("Expected count 3, got: %d", count)
	}
	// Add more items to the set
	set.insert(4)
	set.insert(5)

	// Verify that the count is updated correctly again
	if count := set.Count(); count != 5 {
		t.Errorf("Expected count 5, got: %d", count)
	}
}
