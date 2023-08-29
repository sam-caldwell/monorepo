package keyvalue

import (
	"testing"
)

func TestSetStringHappyPath(t *testing.T) {
	kv := KeyValue{}

	kv.SetString("test_key", "test_value")

	if val, ok := kv.data["test_key"].(string); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != "test_value" {
		t.Errorf("Incorrect value. Expected testvalue, but got %v", val)
	}
}

func TestSetStringOverwriteExistingKey(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"test_key": "old_value",
		},
	}

	kv.SetString("test_key", "new_value")

	if val, ok := kv.data["test_key"].(string); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != "new_value" {
		t.Errorf("Incorrect value. Expected newvalue, but got %v", val)
	}
}

func TestSetStringWithNilMap(t *testing.T) {
	kv := KeyValue{}

	kv.SetString("test_key", "test_value")

	if val, ok := kv.data["test_key"].(string); !ok {
		t.Errorf("Key was not found in map or is not of the correct type")
	} else if val != "test_value" {
		t.Errorf("Incorrect value. Expected testvalue, but got %v", val)
	}
}
