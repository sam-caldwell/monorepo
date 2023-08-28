package argparse

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
	"testing"
)

func TestArguments_Get(t *testing.T) {
	var arg Arguments
	if count := arg.descriptors.Count(); count != 0 {
		t.Fatalf("Expected 0 records. Got %d", count)
	}
	arg.Add("test_get", "-g", "", types.Boolean, true, false, "test get")

	if count := arg.Count(); count != 1 {
		t.Fatalf("Expect 1 record. Got:%d", count)
	}
	if arg.Get("test_get").GetShort() != "-g" {
		t.Fatalf("Expected test_get to have -g short arg")
	}
}
