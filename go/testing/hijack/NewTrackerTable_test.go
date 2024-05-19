package hijack

import (
	"testing"
)

func TestNewTrackerTable(t *testing.T) {
	table := NewTrackerTable()

	// Check if the patches map is initialized
	if table.patches == nil {
		t.Fatalf("Patches map is not initialized")
	}

	// Check if the patches map is empty
	if len(table.patches) != 0 {
		t.Fatalf("Patches map is not empty")
	}
}
