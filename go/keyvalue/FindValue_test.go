package keyvalue

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"testing"
)

func TestFindValue(t *testing.T) {
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
	t.Run("Test: expect value is found", func(t *testing.T) {
		key, found := kv.FindValue(valueExists)
		if !found {
			t.Errorf("Expected value '%s' to be found, but it was not found", valueExists)
		}
		if key != keyExists {
			t.Fatalf("Expected key to be '%s', but it was '%s'", keyExists, key)
		}

	})
	t.Run("Test: expect value is not found", func(t *testing.T) {
		key, found := kv.FindValue(valueNotExists)
		if found {
			t.Errorf("Expected value '%s' to be found, but it was found", valueNotExists)
		}
		if key != words.EmptyString {
			t.Fatalf("Expected key not to be '%s', but it was '%s'", keyNotExists, key)
		}
	})
}
