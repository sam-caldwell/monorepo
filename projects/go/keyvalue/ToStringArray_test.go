package keyvalue

import (
	"testing"
)

func TestToStringArrayHappyPath(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": "value1",
			"key2": "value2",
		},
	}

	output := kv.ToStringArray(" : ", true)

	expectedOutput := []string{
		"key1 : value1",
		"key2 : value2",
	}

	if len(output) != len(expectedOutput) {
		t.Fatalf("Expected %d lines, but got %d", len(expectedOutput), len(output))
	}

	for i, line := range output {
		if line != expectedOutput[i] {
			t.Errorf("Incorrect line. Expected '%s', but got '%s'", expectedOutput[i], line)
		}
	}
}

func TestToStringArrayWithEmptyMap(t *testing.T) {
	kv := KeyValue{}

	output := kv.ToStringArray(" : ", true)

	if len(output) != 0 {
		t.Fatalf("Expected no lines, but got %d", len(output))
	}
}

func TestToStringArrayWithNilMap(t *testing.T) {
	kv := KeyValue{}

	output := kv.ToStringArray(" : ", true)

	if len(output) != 0 {
		t.Fatalf("Expected no lines, but got %d", len(output))
	}
}
