package keyvalue

import (
	"testing"
)

func TestFromCommandShell(t *testing.T) {
	// Initialize a new keyvalue instance
	kv := &KeyValue[string, string]{
		data: make(map[string]string),
	}

	// The command to be tested is `echo "key:value"`
	err := kv.FromCommandShell(":", "\n", "echo", "key:value")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check the key-value pair
	value, ok := kv.data["key"]
	if !ok {
		t.Fatal("Expected key 'key' not found in the keyvalue data")
	}
	if value != "value" {
		t.Fatalf("Unexpected value for 'key': got %v, want 'value'", value)
	}
}
