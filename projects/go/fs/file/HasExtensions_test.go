package file

import (
	"testing"
)

func TestHasExtensions(t *testing.T) {
	t.Run("File with matching extension", func(t *testing.T) {
		filename := "file.yaml"
		extensions := []string{".yaml", ".yml"}
		result := HasExtensions(filename, extensions)

		if !result {
			t.Error("Expected true, but got false")
		}
	})

	t.Run("File with different extension", func(t *testing.T) {
		filename := "file.txt"
		extensions := []string{".yaml", ".yml"}
		result := HasExtensions(filename, extensions)

		if result {
			t.Error("Expected false, but got true")
		}
	})

	t.Run("File with uppercase extension", func(t *testing.T) {
		filename := "file.YML"
		extensions := []string{".yaml", ".yml"}
		result := HasExtensions(filename, extensions)

		if !result {
			t.Error("Expected true, but got false")
		}
	})

	t.Run("File with multiple matching extensions", func(t *testing.T) {
		filename := "file.json"
		extensions := []string{".yaml", ".json"}
		result := HasExtensions(filename, extensions)

		if !result {
			t.Error("Expected true, but got false")
		}
	})

	t.Run("File with no extensions", func(t *testing.T) {
		filename := "file"
		extensions := []string{".yaml", ".yml"}
		result := HasExtensions(filename, extensions)

		if result {
			t.Error("Expected false, but got true")
		}
	})
}
