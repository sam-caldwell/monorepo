package orderedset

import (
	"testing"
)

func TestOrderedSet_Add(t *testing.T) {
	// Create a new set
	testArbitraryTypes := func(data any) {
		set := Set{}
		err := set.Add(42)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
		if set.Count() < 1 {
			t.Errorf("exected a record to exist")
		}
	}
	testArbitraryTypes(1)
	testArbitraryTypes("hi")
	testArbitraryTypes(true)
	testArbitraryTypes(false)
	testArbitraryTypes(1.0)
	testArbitraryTypes(-1)

}
