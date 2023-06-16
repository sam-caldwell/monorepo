package ordered

import (
	"fmt"
	"testing"
)

func TestOrderedSet_Pop(t *testing.T) {
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
	_ = set.Pop()

	// Verify that the set remains unchanged
	if set.Count() != (numberRows - 1) {
		t.Errorf("Expected count %d, got: %d", numberRows-2, set.Count())
	}
}
