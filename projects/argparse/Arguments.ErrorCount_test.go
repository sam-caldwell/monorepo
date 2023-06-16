package argparse

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"testing"
)

func TestArguments_ErrorCount(t *testing.T) {
	var arg Arguments

	if count := arg.descriptors.Count(); count != 0 {
		t.Fatalf("Expected 0 records. Got %d", count)
	}

	//Add an incorrect argument which will push an error to our queue
	arg.Add("a.b", "", "", types.ArgTypes(99), true, nil, "BadRecord")

	if count := arg.ErrorCount(); count != 1 {
		t.Fatal("Expected error did not happen")
	}
}
