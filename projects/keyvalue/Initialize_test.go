package keyvalue

import (
	"testing"
)

// copyData is a helper function to create a copy of the map.
func copyData(dest, src Map) {
	for key, value := range src {
		dest[key] = value
	}
}

func TestInitializeHappyPath(t *testing.T) {
	kv := KeyValue{}
	kv.Initialize(5, true)

	if kv.data == nil {
		t.Error("Failed to initialize keyvalue data")
	}

	// Insert a record
	kv.data["test_key"] = "test_value"
}

func TestInitializeWithoutOverwrite(t *testing.T) {
	kv := KeyValue{}
	kv.Initialize(5, false)

	// Insert a record
	kv.data["test_key"] = "test_value"

	data := make(Map, len(kv.data))
	copyData(data, kv.data)

	kv.Initialize(10, false)

	if len(kv.data) != len(data) {
		t.Error("Initialize should not overwrite existing data when allowOverwrite is set to false")
	}

	// Verify that the data remains unchanged
	for key, value := range data {
		if kv.data[key] != value {
			t.Errorf("Value for key '%s' has changed. Expected: %v, Actual: %v", key, value, kv.data[key])
		}
	}
}

func TestInitializeWithOverwrite(t *testing.T) {
	kv := KeyValue{}
	kv.Initialize(5, false)

	// Insert a record
	kv.data["test_key"] = "test_value"

	if len(kv.data) != 1 {
		t.Errorf("Initialize failed. Expected Size: %d. Actual Size: %d", 1, len(kv.data))
	}

	kv.Initialize(10, true)

	if len(kv.data) != 0 {
		t.Errorf("Initialize failed (overwrite). Expected Size: %d. Actual Size: %d", 0, len(kv.data))
	}
}

func TestInitializeNegativeSize(t *testing.T) {
	kv := KeyValue{}
	kv.Initialize(-5, true)

	if kv.data == nil {
		t.Error("Failed to initialize keyvalue data")
	} else if len(kv.data) != 0 {
		t.Error("Initialize should set size to 0 when negative size is provided")
	}
}
