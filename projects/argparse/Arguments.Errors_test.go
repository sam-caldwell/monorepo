package argparse

import (
	"fmt"
	"testing"
)

func TestArguments_Errors(t *testing.T) {
	// Create an instance of the Arguments struct
	var a = Arguments{}
	_ = a.err.Add(fmt.Errorf("error1"))
	_ = a.err.Add(fmt.Errorf("error2"))
	_ = a.err.Add(fmt.Errorf("error3"))
	_ = a.err.Add(fmt.Errorf("error4"))
	_ = a.err.Add(fmt.Errorf("error5"))

	if a.err.Count() != 5 {
		t.Fatal("Setup failed.  expect 5 errors")
	}
	errorList := a.Errors()

	if len(errorList) != 5 {
		t.Fatal("Expected count 5 not recd")
	}

	for n, e := range errorList {
		expected := fmt.Sprintf("error%d", n+1)
		if e != expected {
			t.Fatalf("%s not found", expected)
		}
	}
}
