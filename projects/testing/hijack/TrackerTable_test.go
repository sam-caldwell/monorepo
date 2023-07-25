package hijack

import (
	"testing"
)

func TestTrackerTableStructure(t *testing.T) {
	table := TrackerTable{}

	if table.patches != nil {
		t.Fatal("patches should be nil")
	}

}
