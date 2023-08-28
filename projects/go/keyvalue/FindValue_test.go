package keyvalue

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/misc/words"
	"testing"
)

func TestFindValue_ValueExists(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": "value1",
			"key2": "value2",
			"key3": "value1",
		},
	}

	key, found := kv.FindValue("value2")

	if !found {
		t.Errorf("Expected value 'value2' to be found, but it was not found")
	}

	if key != "key2" {
		t.Errorf("Expected key 'key2', but got '%v'", key)
	}
}

func TestFindValue_ValueDoesNotExist(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": "value1",
			"key2": "value2",
		},
	}

	key, found := kv.FindValue("value3")

	if found {
		t.Errorf("Expected value 'value3' to not be found, but it was found")
	}

	if key != words.EmptyString {
		t.Errorf("Expected key to be an empty string, but got '%v'", key)
	}
}
