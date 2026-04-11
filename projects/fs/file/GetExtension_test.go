package file

import (
	"testing"
)

func TestGetExtension(t *testing.T) {
	t.Run("File with extension", func(t *testing.T) {
		filename := "file.txt"
		extension := GetExtension(filename)
		expected := ".txt"

		if extension != expected {
			t.Fatalf("Expected extension '%s', but got '%s'", expected, extension)
		}
	})

	t.Run("File without extension", func(t *testing.T) {
		filename := "file"
		extension := GetExtension(filename)
		expected := ""

		if extension != expected {
			t.Fatalf("Expected extension '%s', but got '%s'", expected, extension)
		}
	})

	t.Run("File with uppercase extension", func(t *testing.T) {
		filename := "file.TXT"
		extension := GetExtension(filename)
		expected := ".txt"

		if extension != expected {
			t.Fatalf("Expected extension '%s', but got '%s'", expected, extension)
		}
	})

	t.Run("File with multiple extensions", func(t *testing.T) {
		filename := "file.tar.gz"
		extension := GetExtension(filename)
		expected := ".gz"

		if extension != expected {
			t.Fatalf("Expected extension '%s', but got '%s'", expected, extension)
		}
	})
}
