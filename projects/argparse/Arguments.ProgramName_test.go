package argparse

import (
	"github.com/sam-caldwell/go/v2/projects/wrappers/os"
	"path"
	"testing"
)

func TestArguments_ProgramName(t *testing.T) {
	var arg Arguments

	expected := path.Base(os.Args[0])

	if actual := arg.ProgramName().programName; actual != expected {
		t.Fatalf("Program name mismatch\n%s\n%s", actual, expected)
	}
}
