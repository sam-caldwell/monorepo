package keyvalue

import (
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
		t.Run("initialize data", func(t *testing.T) {
			kv.data = map[string]string{
				"key1": "value1",
				"key2": "value2",
			}
		})
		t.Run("verify result", func(t *testing.T) {
			expectedOutput := "key1:value1\nkey2:value2\n"
			actualOutput := kv.ToString(":", "\n")
			if actualOutput != expectedOutput {
				t.Fatalf("Incorrect output. Expected:\n'%s'\n\nBut got:\n'%s'\n", expectedOutput, actualOutput)
			}
		})
	})
}
