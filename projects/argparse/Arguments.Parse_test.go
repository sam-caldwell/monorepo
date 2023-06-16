package argparse

import (
	"github.com/sam-caldwell/go/v2/projects/argparse/types"
	"os"
	"testing"
)

func TestArguments_Parse(t *testing.T) {
	var args Arguments

	//failIf := func(condition bool, msg string, params ...any) {
	//	if condition {
	//		t.Fatalf(msg, params)
	//	}
	//}

	os.Args = []string{
		"argparse.test", //Program name
		"pos1value",     //positional arg 0
		"pos2value",     //positional arg 1
		"-a",            //short arg (flag)
		"-n", "5",       //short arg (int)
		"-b", "true", //short arg (boolean)
		"-f", "1.0", //short arg (float)
		"--long", "yes", //long arg (string)
	}

	args.
		ProgramName().
		Copyright(2023, "Sam Caldwell", "mail@samcaldwell.net>").
		Preamble("This is a description of our Arguments Object.").
		Postscript("This is the postfix string after help is dumped.").
		Add("a0pos", "", "", types.String, true, "not set", "pos  1").
		Add("a1pos", "", "", types.String, true, "not set", "pos  2").
		Add("all", "-a", "--all", types.Flag, true, false, "All flag").
		Add("number", "-n", "--number", types.Integer, true, 0, "set num").
		Add("bool", "-b", "--bool", types.Boolean, true, false, "set bool").
		Add("long", "-l", "--long", types.String, true, "no", "set string").
		Add("float", "-f", "--float", types.Float, true, 3.14, "pi").
		Parse().
		ExitOnError()
	//
	//if args.ErrorCount() != 0 {
	//	for _, e := range args.Errors() {
	//		t.Logf("Error: %s", e)
	//	}
	//	t.Fatalf("Errors encountered (count: %d)", args.ErrorCount())
	//}
}
