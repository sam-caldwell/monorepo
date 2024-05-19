package keyvalue

import (
	"strings"
	"testing"
)

func TestKeyValue_ToString(t *testing.T) {
	t.Run("test with nil data", func(t *testing.T) {
		var kv KeyValue[string, string]
		output := kv.ToString(":", "\n")
		if output != "" {
			t.Fatalf("Incorrect output. Expected empty string, but got:'%s'", output)
		}
	})
	t.Run("test with data", func(t *testing.T) {
		var kv KeyValue[string, string]
		var expectedOutput string
		t.Run("initialize data", func(t *testing.T) {
			kv.data = map[string]string{
				"key1": "value1",
				"key2": "value2",
			}
			expectedOutput = strings.Join([]string{
				"key1:value1",
				"key2:value2",
			}, "\n")
		})
		t.Run("verify result", func(t *testing.T) {
			output := kv.ToString(":", "\n")
			if output != expectedOutput {
				t.Fatalf("Incorrect output. Expected:\n%s\n\nBut got:\n%s", expectedOutput, output)
			}
		})
	})
}
