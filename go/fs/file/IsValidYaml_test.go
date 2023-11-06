package file

import (
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"testing"
)

func TestIsValidYaml(t *testing.T) {
	const testFile = "testfile.yaml"
	const testFilePerms = 0644
	const yamlContent = `
key1: value1
key2: value2
`

	const invalidYamlContent = `
key1: value1
key2
`
	/*
	 *
	 */
	t.Run("Valid YAML file", func(t *testing.T) {
		// Create a temporary YAML file for testing

		tempFile, err := os.CreateTemp("", testFile)
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer func() {
			// Clean up the temporary file after the test
			if err = os.Remove(tempFile.Name()); err != nil {
				t.Errorf("Failed to remove the temporary file: %v", err)
			}
		}()

		// Write the YAML content to the temporary file
		err = os.WriteFile(tempFile.Name(), []byte(yamlContent), testFilePerms)
		if err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}

		err = IsValidYaml(tempFile.Name())
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
	})
	/*
	 *
	 */
	t.Run("Invalid YAML file", func(t *testing.T) {
		// Create a temporary file with invalid YAML content for testing

		tempFile, err := os.CreateTemp("", testFile)
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer func() {
			// Clean up the temporary file after the test
			if err = os.Remove(tempFile.Name()); err != nil {
				t.Errorf("Failed to remove the temporary file: %v", err)
			}
		}()

		// Write the invalid YAML content to the temporary file
		err = os.WriteFile(tempFile.Name(), []byte(invalidYamlContent), testFilePerms)
		if err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}

		err = IsValidYaml(tempFile.Name())
		if err == nil {
			t.Error("Expected an error, but got nil")
		}
	})
}
