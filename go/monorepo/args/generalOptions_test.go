package args

import (
	"testing"
)

func TestArguments_generalOptions(t *testing.T) {
	testCase := func(name string, args []string, remainingCount int, debug, color, noop bool) {
		arg := &Arguments{}
		//args := []string{"program", "--debug", "--color", "--noop", "arg2"}
		remainingArgs := arg.generalOptions(args)
		if arg.option.color != color {
			t.Fatalf("(case: %s) Expected color flag to be %v, got %v", name, color, arg.option.color)
		}
		if arg.option.debug != debug {
			t.Fatalf("(case: %s) Expected debug flag to be %v, got %v", name, debug, arg.option.debug)
		}
		if arg.option.noop != noop {
			t.Fatalf("(case: %s) Expected noop flag to be %v, got %v", name, noop, arg.option.noop)
		}
		if len(remainingArgs) != remainingCount {
			t.Fatalf("(case: %s) Expected %d remaining arguments, got %d [%v]",
				name, remainingCount, len(remainingArgs), remainingArgs)
		}
		//t.Logf("Test Case (%s): OK", name)
	}
	//
	// No General Options present
	testCase("1", []string{"program", "command", "arg1", "arg2"}, 4, false, false, false)
	//
	// The General Options are present but not first (and should not be consumed)
	testCase("2", []string{"program", "--debug", "--color", "--noop", "arg1"}, 5, false, false, false)
	//
	// The General Options are present AND first (and should be consumed)
	testCase("3", []string{"--debug", "--color", "--noop", "arg1", "arg2"}, 2, true, true, true)
	//
	// The General Option --debug is present and first but a non-relevant option is between others.
	testCase("4", []string{"--debug", "arg1", "--color", "--noop", "arg2"}, 4, true, false, false)
	//
	// The General Option --debug is present and first but a non-relevant option is between others.
	testCase("5", []string{"--debug", "--color", "--noop"}, 0, true, true, true)

}
