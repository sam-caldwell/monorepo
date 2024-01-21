package file

import (
	"fmt"
	"testing"
)

func TestFile_SetWithExtension(t *testing.T) {

	testFunc := func(index int, fileName string, extensions []string, expectedError error) {
		var f File
		if string(f) != "" {
			t.Fatal("expect initial empty value")
		}
		err := f.SetWithExtension(fileName, extensions)
		if (err != nil) && (expectedError != nil) && (err.Error() != expectedError.Error()) {
			t.Fatalf("%d error mismatch. file '%s' err '%v' expectedError '%v'",
				index, fileName, err, expectedError)
		}
	}

	type TestData struct {
		fileName   string
		extensions []string
		error      error
	}
	testData := []TestData{
		{"testFile.iso", []string{"iso"}, nil},
		{"testFile.iso", []string{".iso"}, nil},
		{"testFile.iso", []string{".iso", ".img", ".disk"}, nil},
		{"testFile.img", []string{".iso", ".img", ".disk"}, nil},
		{"testFile.disk", []string{".iso", ".img", ".disk"}, nil},
		{"testFile", []string{".pdf", ".doc"}, nil},
		{"testFile", []string{"iso", ".iso"}, fmt.Errorf("missing extension")},
		{"testFile", []string{".pdf", ".doc"}, fmt.Errorf("missing extension")},
	}

	for index, row := range testData {
		testFunc(index, row.fileName, row.extensions, row.error)
	}
}
