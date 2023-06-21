package file

import (
	"testing"
)

func TestHasExtension(t *testing.T) {
	t.Run("File with matching extension", func(t *testing.T) {
		filename := "file.yaml"
		extension := ".yaml"
		result := HasExtension(filename, extension)

		if !result {
			t.Error("Expected true, but got false")
		}
	})

	t.Run("File with different extension", func(t *testing.T) {
		filename := "file.txt"
		extension := ".yaml"
		result := HasExtension(filename, extension)

		if result {
			t.Error("Expected false, but got true")
		}
	})

	t.Run("File with uppercase extension", func(t *testing.T) {
		filename := "file.YAML"
		extension := ".yaml"
		result := HasExtension(filename, extension)

		if !result {
			t.Error("Expected true, but got false")
		}
	})
}
