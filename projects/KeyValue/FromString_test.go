package keyvalue

import (
	"reflect"
	"testing"
)

func TestFromString(t *testing.T) {
	kv := KeyValue{}

	data := "key1:value1\nkey2:value2\nkey3:value3"
	lineEnding := "\n"
	columnDelimiter := ":"

	kv.FromString(&data, lineEnding, columnDelimiter)

	expectedData := Map{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	if !reflect.DeepEqual(kv.data, expectedData) {
		t.Errorf("Parsed data does not match expected data")
	}
}
