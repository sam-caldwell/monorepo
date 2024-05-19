package keyvalue

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"testing"
)

func TestFindKey(t *testing.T) {
	const (
		keyExists      = "key3"
		valueExists    = "value3"
		keyNotExists   = "key9"
		valueNotExists = "value9"
	)
	kv := KeyValue[string, string]{
		data: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
			"key4": "value4",
			"key5": "value5",
			"key6": "value6",
		},
	}
	t.Run("Test: expect key is found", func(t *testing.T) {
		value, found := kv.FindKey(keyExists)
		if !found {
			t.Fatalf("Expected key '%s' to be found, but it was not found", keyExists)
		}
		if value != valueExists {
			t.Fatalf("Expected value '%s', but got '%s'", valueExists, value)
		}
	})
	t.Run("Test: expect key is not found", func(t *testing.T) {
		value, found := kv.FindKey(keyNotExists)
		if found {
			t.Fatalf("Expected key '%s' not to be found, but it was found", keyNotExists)
		}
		if value != words.EmptyString {
			t.Fatalf("Expected value '%s', but got '%s'", valueNotExists, value)
		}
	})
}
