package keyvalue

import (
	"os"
	"testing"
)

func TestKeyValueFromFile(t *testing.T) {
	const (
		testFile = "testFile"
		tempDir  = "/tmp"
	)

	var (
		kv       KeyValue[string, string]
		fileName string
	)

	defer func() { _ = os.Remove(fileName) }()
	t.Run("Create temp source file", func(t *testing.T) {
		content := []byte("#comment:line\nkey1:value1 #comment1\nkey2:value2\nkey3:value3")
		tempFile, err := os.CreateTemp(tempDir, testFile)
		if err != nil {
			t.Fatalf("Failed to create temporary file: %s", err)
		}
		defer func() { _ = tempFile.Close() }()
		fileName = tempFile.Name()

		// Write test content to the temporary file
		if _, err := tempFile.Write(content); err != nil {
			t.Fatalf("Failed to write to temporary file: %s", err)
		}
	})
	t.Run("Test import of file content to KeyValue", func(t *testing.T) {
		if err := kv.FromFile(fileName, ':', '\n', '#'); err != nil {
			t.Errorf("FromFile returned an error: %s", err)
		}
	})
	t.Run("Confirm the KeyValue (kv) data is valid", func(t *testing.T) {
		t.Run("confirm key1", func(t *testing.T) {
			expectedValue := "value1"
			actualValue, exists := kv.FindKey("key1")
			if !exists {
				t.Errorf("Expected key1 to exist in keyvalue")
			}
			if actualValue != expectedValue {
				t.Errorf("Expected value for key1: %s, but got: %s", expectedValue, actualValue)
			}
		})
		t.Run("confirm comment is not found", func(t *testing.T) {
			_, e := kv.FindKey("comment")
			if e {
				t.Fatal("comment should not be found")
			}
			_, e = kv.FindKey("#comment")
			if e {
				t.Fatal("comment should not be found")
			}
		})
		t.Run("confirm all keys in set", func(t *testing.T) {
			expected := map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			}
			for key, value := range expected {
				actualValue, exists := kv.FindKey(key)
				if !exists {
					t.Errorf("Expected %s to exist in keyvalue", key)
				}
				if actualValue != value {
					t.Errorf("Expected value for %s: %s, but got: %s", key, value, actualValue)
				}
			}
		})
	})

}
