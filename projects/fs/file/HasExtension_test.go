package file

import (
	"testing"
)

func TestHasExtension(t *testing.T) {
	testFunc := func(index int, fileName, extension string, expectedResult bool) {
		t.Run("File with matching extension", func(t *testing.T) {
			if actualResult := HasExtension(fileName, extension); actualResult != expectedResult {
				t.Fatalf("%d Expected %v, but got %v on '%s' with '%v'",
					index, expectedResult, actualResult, fileName, extension)
			}
		})
	}
	type TestData struct {
		name      string
		extension string
		result    bool
	}
	testData := []TestData{
		{"file.yaml", ".yaml", true},
		{"file.yaml", ".YAML", true},
		{"file.json", ".json", true},
		{"file.yml", ".yml", true},
		{"file.iso", ".iso", true},
		{"file.iso", "iso", false},
		{"file.yml", "yml", false},
		{"file.yaml", "yaml", false},
		{"file.yaml", "YAML", false},
		{"file.json", "json", false},
		{"file", "yaml", false},
		{"file", "yml", false},
		{"file", "json", false},
		{"file", "iso", false},
		{"file", "", true},
	}
	for n, row := range testData {
		testFunc(n, row.name, row.extension, row.result)
	}
}
