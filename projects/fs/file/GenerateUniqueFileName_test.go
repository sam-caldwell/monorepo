package file

import (
	"strings"
	"testing"
)

func TestGenerateUniqueFileName(t *testing.T) {
	t.Run("Generate file with no collision", func(t *testing.T) {
		actual, err := GenerateUniqueFileName("/tmp", ".tql", 3)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.HasSuffix(actual, ".tql") {
			t.Fatal("expected .tql file extension")
		}
		if Exists(actual) {
			t.Fatal("file should not exist.  Only a file name should be created.")
		}
	})
}
