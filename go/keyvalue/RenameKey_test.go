package keyvalue

import "testing"

func TestKeyValue_RenameKey(t *testing.T) {
	var kv KeyValue[string, int]
	t.Run("initialize data", func(t *testing.T) {
		kv.data = map[string]int{
			"oldKey": 1,
			"key2":   2,
		}
	})
	t.Run("rename key to newKey", func(t *testing.T) {
		if !kv.RenameKey("oldKey", "newKey") {
			t.Fatalf("Failed to rename existing key")
		}
		if _, ok := kv.data["oldKey"]; ok {
			t.Fatalf("Old key still exists after renaming")
		}
		if _, ok := kv.data["newKey"]; !ok {
			t.Fatalf("New key does not exist after renaming")
		}
	})
	t.Run("Rename non-existing key", func(t *testing.T) {
		if kv.RenameKey("nonExistent", "shouldNotExist") {
			t.Fatalf("RenameKey should return false when trying to rename a non-existent key")
		}
	})
	t.Run("Rename nil key", func(t *testing.T) {
		var nilKv KeyValue[string, int]
		if nilKv.data != nil {
			t.Fatalf("expect nil data map")
		}
		if nilKv.RenameKey("ShouldNotExist", "willNotExist") {
			t.Fatalf("RenameKey() should return false on nil map")
		}
	})
}
