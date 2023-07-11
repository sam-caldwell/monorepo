package simpleArgs

import (
	"github.com/sam-caldwell/go/v2/projects/wrappers/os"
	"testing"
)

func TestHasFlag(t *testing.T) {

	data := map[string]any{
		"--test0": nil,
		"--test1": nil,
		"--test2": "value2",
		"--test3": "value3",
		"--test4": "value4",
		"--test5": nil,
		"--test6": nil,
		"--test7": "value7",
		"--test8": nil,
	}

	//Setup our test data
	for key, value := range data {
		os.Args = append(os.Args, key)
		if value != nil {
			os.Args = append(os.Args, value.(string))
		}
	}

	for expectedKey, _ := range data {
		if !HasFlag(expectedKey) {
			t.Fatalf("expected key not found")
		}
	}
	if HasFlag("--fictionalNonExistentKey") {
		t.Fatal("non-existent key not found")
	}
}
