package simple

import (
	"testing"
)

func TestSet_Has(t *testing.T) {
	// Create a new set
	set := Set{}

	// Check if the set has no items initially
	if set.Has("apple") {
		t.Fatal("Expected 'apple' to not exist in the set, but it does")
	}

	if set.Has(true) {
		t.Fatal("Expected true to not exist in the set, but it does")
	}

	if set.Has(42) {
		t.Fatal("Expected 42 to not exist in the set, but it does")
	}

	// Add items of different types to the set
	if err := set.Add("apple"); err != nil {
		t.Fatal(err)
	}
	if !set.Has("apple") {
		t.Fatal("expected 'apple' to exist")
	}
	if set.Has("not found") {
		t.Fatal("expected 'not found' to not exist but it does.")
	}
}
