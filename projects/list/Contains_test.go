package list

import "testing"

func TestContains(t *testing.T) {
	// Happy path scenario
	sourceList := []string{"apple", "banana", "orange"}
	value := "banana"
	result := Contains(sourceList, value)
	if !result {
		t.Fatalf("Expected %s to be in the list, but it's not", value)
	}

	// Sad path scenario
	value = "grape"
	result = Contains(sourceList, value)
	if result {
		t.Fatalf("Expected %s not to be in the list, but it is", value)
	}

	// Additional happy path scenario
	value = "apple"
	result = Contains(sourceList, value)
	if !result {
		t.Fatalf("Expected %s to be in the list, but it's not", value)
	}
}
