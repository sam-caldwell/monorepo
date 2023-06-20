package keyvalue

import (
	"strings"
	"testing"
)

func TestToStringHappyPath(t *testing.T) {
	kv := KeyValue{
		data: Map{
			"key1": "value1",
			"key2": "value2",
		},
	}

	output := kv.ToString(" : ", "\n")

	expectedOutput := strings.Join([]string{
		"key1 : value1",
		"key2 : value2",
	}, "\n")

	if output != expectedOutput {
		t.Errorf("Incorrect output. Expected:\n%s\n\nBut got:\n%s", expectedOutput, output)
	}
}

func TestToStringWithEmptyMap(t *testing.T) {
	kv := KeyValue{}

	output := kv.ToString(" : ", "\n")

	if output != "" {
		t.Errorf("Incorrect output. Expected empty string, but got:\n%s", output)
	}
}

func TestToStringWithNilMap(t *testing.T) {
	kv := KeyValue{}

	output := kv.ToString(" : ", "\n")

	if output != "" {
		t.Errorf("Incorrect output. Expected empty string, but got:\n%s", output)
	}
}
