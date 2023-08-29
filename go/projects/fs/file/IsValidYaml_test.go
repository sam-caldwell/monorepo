package file

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/wrappers/os"
	"io/ioutil"
	"testing"
)

func TestIsValidYaml(t *testing.T) {
	t.Run("Valid YAML file", func(t *testing.T) {
		// Create a temporary YAML file for testing
		yamlContent := `
key1: value1
key2: value2
`
		tempFile, err := ioutil.TempFile("", "testfile.yaml")
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
		err = ioutil.WriteFile(tempFile.Name(), []byte(yamlContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}

		err = IsValidYaml(tempFile.Name())
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
	})

	t.Run("Invalid YAML file", func(t *testing.T) {
		// Create a temporary file with invalid YAML content for testing
		invalidYamlContent := `
key1: value1
key2
`

		tempFile, err := ioutil.TempFile("", "testfile.yaml")
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
		err = ioutil.WriteFile(tempFile.Name(), []byte(invalidYamlContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}

		err = IsValidYaml(tempFile.Name())
		if err == nil {
			t.Error("Expected an error, but got nil")
		}
	})
}
