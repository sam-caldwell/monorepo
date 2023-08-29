package keyvalue

import (
	"testing"
)

func TestRenameKeyHappyPath(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"oldkey": 1,
			"key2":   2,
		},
	}

	if !kv.RenameKey("oldkey", "newkey") {
		t.Errorf("Failed to rename existing key")
	} else if _, ok := kv.data["newkey"]; !ok {
		t.Errorf("New key does not exist after renaming")
	} else if _, ok := kv.data["oldkey"]; ok {
		t.Errorf("Old key still exists after renaming")
	}
}

func TestRenameKeyNonExistentKey(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": 1,
			"key2": 2,
		},
	}

	if kv.RenameKey("nonexistent", "newkey") {
		t.Errorf("RenameKey should return false when trying to rename a non-existent key")
	}
}

func TestRenameKeyWithNilMap(t *testing.T) {
	kv := KeyValue{}

	if kv.RenameKey("oldkey", "newkey") {
		t.Errorf("RenameKey should return false when called on a keyvalue struct with a nil map")
	}
}
