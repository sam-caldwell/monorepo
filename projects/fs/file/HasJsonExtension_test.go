package file

import (
	"testing"
)

func TestHasJsonExtension(t *testing.T) {
	t.Run("File with matching extension", func(t *testing.T) {
		filename := "file.json"
		result := HasJsonExtension(filename)

		if !result {
			t.Error("Expected true, but got false")
		}
	})

	t.Run("File with different extension", func(t *testing.T) {
		filename := "file.txt"
		result := HasJsonExtension(filename)

		if result {
			t.Error("Expected false, but got true")
		}
	})

	t.Run("File with uppercase extension", func(t *testing.T) {
		filename := "file.JSON"
		result := HasJsonExtension(filename)

		if !result {
			t.Error("Expected true, but got false")
		}
	})
}
