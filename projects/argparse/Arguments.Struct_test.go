package argparse

import (
	"testing"
)

func TestArguments_Struct(t *testing.T) {
	var arg Arguments

	if arg.pos.Value() != 0 {
		t.Fail()
	}
	if arg.descriptors.Count() != 0 {
		t.Fail()
	}
	if arg.programName != "" {
		t.Fail()
	}
	if arg.preambleText.Count() != 0 {
		t.Fail()
	}
	if arg.postscriptText.Count() != 0 {
		t.Fail()
	}
	if arg.err.Count() != 0 {
		t.Fail()
	}
	if arg.results.Count() != 0 {
		t.Fail()
	}

}
