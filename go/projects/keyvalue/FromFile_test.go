package keyvalue

import (
	os2 "github.com/sam-caldwell/monorepo/go/projects/v2/wrappers/os"
	"testing"
)

func TestKeyValueFromFile(t *testing.T) {
	const (
		testFile = "testfile"
		tempDir  = "/tmp"
	)
	// Create a temporary file for testing
	content := []byte("#comment:line\nkey1:value1 #comment1\nkey2:value2\nkey3:value3")
	tempFile, err := os2.CreateTemp(tempDir, testFile)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %s", err)
	}
	defer func() { _ = os2.Remove(tempFile.Name()) }()

	// Write test content to the temporary file
	if _, err := tempFile.Write(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %s", err)
	}

	// Close the file before passing it to the FromFile function
	_ = tempFile.Close()

	// Create a new instance of keyvalue
	kv := &KeyValue{}

	// Call FromFile method
	err = kv.FromFile(tempFile.Name(), ":", "\n")
	if err != nil {
		t.Errorf("FromFile returned an error: %s", err)
	}

	// Perform assertions on the keyvalue struct
	{
		expectedValue := "value1"
		actualValue, exists := kv.FindKey("key1")
		if !exists {
			t.Errorf("Expected key1 to exist in keyvalue")
		}
		if actualValue != expectedValue {
			t.Errorf("Expected value for key1: %s, but got: %s", expectedValue, actualValue)
		}
	}
	{
		_, e := kv.FindKey("comment")
		if e {
			t.Fatal("comment should not be found")
		}
		_, e = kv.FindKey("#comment")
		if e {
			t.Fatal("comment should not be found")
		}
	}

	{
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
	}

}
