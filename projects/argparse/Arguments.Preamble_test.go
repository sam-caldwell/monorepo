package argparse

import (
	"fmt"
	"testing"
)

func TestArguments_Preamble(t *testing.T) {
	var arg Arguments

	arg.Preamble("test0").
		Preamble("test1").
		Preamble("test2")

	for i := range arg.preambleText.List() {
		expected := fmt.Sprintf("test%d", i)
		if line := arg.preambleText.Pop(); line != expected {
			t.Fatalf("%s expected Got '%s'", expected, line)
		}
	}

}
