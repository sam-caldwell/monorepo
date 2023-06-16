package simple

import (
	"fmt"
	"testing"
)

func TestSet_Delete(t *testing.T) {
	// Create a new set
	set := Set{}

	const numberRows = 4
	for i := 1; i <= numberRows; i++ {
		// Add items to the set
		if err := set.Add(fmt.Sprintf("test%d", i)); err != nil {
			t.Fatal(err)
		}
	}
	if set.Count() != numberRows {
		t.Errorf("Expected count %d, got: %d", numberRows, set.Count())
	}

	// Delete an existing item
	if err := set.Delete("test1"); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Verify that the item is removed from the set
	if set.Has("test1") {
		t.Error("Expected 'test1' to be deleted from the set, but it still exists")
	}
	// Verify that the item is removed from the set
	if !set.Has("test2") {
		t.Error("Expected 'test2' should exist")
	}
	// Verify that the item is removed from the set
	if !set.Has("test3") {
		t.Error("Expected 'test3' should exist")
	}

	// Verify that the set remains unchanged
	if set.Count() != (numberRows - 1) {
		t.Errorf("Expected count %d, got: %d", numberRows-1, set.Count())
	}
}
