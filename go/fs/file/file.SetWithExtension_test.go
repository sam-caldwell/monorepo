package file

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
	"testing"
)

func TestFile_SetWithExtension(t *testing.T) {

	testFunc := func(index int, fileName string, extensions []string, expectedError error) {
		t.Run(fmt.Sprintf("test %d (name:%s, extensions:[%v],expectedError:%v)",
			index, fileName, strings.Join(extensions, words.Comma), expectedError),
			func(t *testing.T) {

				var f File

				ansi.Cyan().Printf("test with '%s' and %v ", fileName, extensions).Reset()

				err := f.ValidIfHasExtension(fileName, extensions)

				if expectedError == nil {
					if err != nil {
						ansi.Red().Printf("fail").LF().Reset()
						t.Fatalf("%d error mismatch. file '%s' err '%v' expect no error", index, fileName, err)
					} else {
						ansi.Green().Printf("ok").LF().Reset()
						return
					}
				}
				if (err != nil) && (err.Error() != expectedError.Error()) {
					ansi.Red().Printf("fail").LF().Reset()
					t.Fatalf("%d error mismatch. file '%s' err '%v' expectedError '%v'",
						index, fileName, err, expectedError)
				}
				if err == nil {
					ansi.Red().Printf("fail").LF().Reset()
					t.Fatalf("%d expected error not found. file '%s' expectedError '%v'",
						index, fileName, expectedError)
				}
				ansi.Green().Printf("ok").LF().Reset()
			})
	}

	type TestData struct {
		fileName   string
		extensions []string
		error      error
	}
	testData := []TestData{
		{"testFile.iso", []string{".iso"}, nil},
		{"testFile.iso", []string{".iso", ".img", ".disk"}, nil},
		{"testFile.img", []string{".iso", ".img", ".disk"}, nil},
		{"testFile.disk", []string{".iso", ".img", ".disk"}, nil},
		{"testFile.pdf", []string{".pdf", ".doc"}, nil},
		{"testFile.doc", []string{".pdf", ".doc"}, nil},
		{"testFile.iso", []string{".pdf", ".doc"}, fmt.Errorf("missing extension")},
		{"testFile", []string{".img", ".iso"}, fmt.Errorf("missing extension")},
		{"testFile", []string{".pdf", ".doc"}, fmt.Errorf("missing extension")},
		{"test", []string{".iso"}, fmt.Errorf("missing extension")},
	}

	for index, row := range testData {
		testFunc(index, row.fileName, row.extensions, row.error)
	}
}
