package hijack

import (
	"reflect"
	"testing"
)

func TestTrackerTableStructure(t *testing.T) {
	table := TrackerTable{}

	// Check if the table has a lock of type sync.Mutex
	if _, ok := reflect.TypeOf(table.lock).FieldByName("Mutex"); !ok {
		t.Errorf("TrackerTable lock is not of type sync.Mutex")
	}

	// Check if the patches field is of type map[uintptr]AppliedPatch
	if _, ok := reflect.TypeOf(table.patches).FieldByName("patches"); !ok {
		t.Errorf("TrackerTable patches field is not of type map[uintptr]AppliedPatch")
	}
}
