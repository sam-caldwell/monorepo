package hijack

import (
	"testing"
)

func TestGlobalPatchTable_Delete(t *testing.T) {
	// Create a sample record to delete
	recordKey := uintptr(123)
	recordValue := AppliedPatch{}

	// Create a TrackerTable instance with the sample record
	table := TrackerTable{
		patches: map[uintptr]AppliedPatch{
			recordKey: recordValue,
		},
	}

	// Call the delete method to delete the record
	err := table.delete(recordKey)

	// Check if an error occurred
	if err != nil {
		t.Errorf("Unexpected error while deleting record: %v", err)
	}

	// Check if the record is deleted from the map
	_, exists := table.patches[recordKey]
	if exists {
		t.Errorf("Record with key %v still exists in the map", recordKey)
	}
}
