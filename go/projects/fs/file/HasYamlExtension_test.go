package file

import (
	"testing"
)

func TestHasYamlExtension(t *testing.T) {
	t.Run("File with long YAML extension", func(t *testing.T) {
		filename := "file.yaml"
		result := HasYamlExtension(filename)

		if !result {
			t.Error("Expected true, but got false")
		}
	})

	t.Run("File with short YAML extension", func(t *testing.T) {
		filename := "file.yml"
		result := HasYamlExtension(filename)

		if !result {
			t.Error("Expected true, but got false")
		}
	})

	t.Run("File with different extension", func(t *testing.T) {
		filename := "file.txt"
		result := HasYamlExtension(filename)

		if result {
			t.Error("Expected false, but got true")
		}
	})

	t.Run("File with uppercase extension", func(t *testing.T) {
		filename := "file.YML"
		result := HasYamlExtension(filename)

		if !result {
			t.Error("Expected true, but got false")
		}
	})
}
