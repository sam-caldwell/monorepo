package simple

import "testing"

func TestSet_Count(t *testing.T) {
	t.Skip("debugging")
	//Create a new set
	var set Set[int]

	// Verify that the initial count is 0
	if count := set.Count(); count != 0 {
		t.Errorf("Expected count 0, got: %d", count)
	}

	// Add items to the set
	if err := set.Add(1); err != nil {
		t.Fatal(err)
	}
	if err := set.Add(2); err != nil {
		t.Fatal(err)
	}
	if err := set.Add(3); err != nil {
		t.Fatal(err)
	}

	// Verify that the count is updated correctly
	if count := set.Count(); count != 3 {
		t.Errorf("Expected count 3, got: %d", count)
	}

	// Add more items to the set
	if err := set.Add(4); err != nil {
		t.Fatal(err)
	}
	if err := set.Add(5); err != nil {
		t.Fatal(err)
	}

	// Verify that the count is updated correctly again
	if count := set.Count(); count != 5 {
		t.Errorf("Expected count 5, got: %d", count)
	}
}
