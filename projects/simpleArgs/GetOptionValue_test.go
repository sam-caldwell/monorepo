package simpleArgs

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/wrappers/os"
	"testing"
)

func TestGetOptionValue(t *testing.T) {
	type TestRecord struct {
		value any
		err   error
	}
	optionHasNoValue := fmt.Errorf("option has no value")

	data := map[string]TestRecord{
		"--test0": {value: nil, err: optionHasNoValue},
		"--test1": {value: nil, err: optionHasNoValue},
		"--test2": {value: "value2", err: nil},
		"--test3": {value: "value3", err: nil},
		"--test4": {value: "value4", err: nil},
		"--test5": {value: nil, err: optionHasNoValue},
		"--test6": {value: nil, err: optionHasNoValue},
		"--test7": {value: "value7", err: nil},
		"--test8": {value: nil, err: optionHasNoValue},
	}

	//Setup our test data
	for key, record := range data {
		os.Args = append(os.Args, key)
		if record.value != nil {
			os.Args = append(os.Args, record.value.(string))
		}
	}

	for expectedKey, expected := range data {
		actualValue, err := GetOptionValue(expectedKey)
		if expected.value == nil {
			if err == nil {
				t.Fatalf("Expected error on %s but got none", expectedKey)
			}
			if err.Error() != expected.err.Error() {
				t.Fatalf("error mismatch\n"+
					"  actual:%v\n"+
					"expected:%v", err.Error(), expected.err.Error())
			}
		} else {
			if actualValue != expected.value {
				t.Fatalf("value mismatch\n"+
					"  actual:%v\n"+
					"expected:%v", actualValue, expected.value)
			}
		}
	}
}
