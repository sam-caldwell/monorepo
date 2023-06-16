package argparse

import (
	"github.com/sam-caldwell/go-monorepo/v2/projects/argparse/types"
	"strings"
	"testing"
)

const expected = `argparse.test
  This is a test program
  It does things...really cool things.

Usage:
  argparse.test [positional args] [optional args]

  Positional Arguments
    noshort5 [String] - test no5
    pos0 [Integer] - positional0

 Optional Arguments
                 --nos5  [String]  [default:ns]     - test no5              
    -1 [Boolean] --test1 [Boolean] [default:true]   - Test 1                
    -2 [Boolean] --all2  [Boolean] [default:true]   - Test all2             
    -3 [String]  --now3  [String]  [default:notnow] - test now3             
    -4           --now4                             - test now4             
    -h           --help                             - show this help message

 This program has a postscript
 The postscript comes after usage.
 (c) 2023 Sam Caldwell <mail@samcaldwell.net>`

func TestArguments_Help(t *testing.T) {
	var arg Arguments
	arg.ProgramName().
		Preamble("This is a test program").
		Preamble("It does things...really cool things.").
		Postscript("This program has a postscript").
		Postscript("The postscript comes after usage.").
		Copyright(2023, "Sam Caldwell", "mail@samcaldwell.net").
		Add("pos0", "", "", types.Integer, true, 1, "positional0").
		Add("test1", "-1", "--test1", types.Boolean, true, true, "Test 1").
		Add("all2", "-2", "--all2", types.Boolean, true, true, "Test all2").
		Add("now3", "-3", "--now3", types.String, true, "notnow", "test now3").
		Add("now4", "-4", "--now4", types.Flag, true, false, "test now4").
		Add("noshort5", "", "--nos5", types.String, true, "ns", "test no5")

	helpText := arg.Help()

	if arg.HasErrors() {
		t.Log("Errors:")
		for _, e := range arg.Errors() {
			t.Log(e)
		}
		t.Log("----")
		t.Fatal("Failing due to errors")
	}

	actualText := strings.TrimSpace(helpText)
	expectedText := strings.TrimSpace(expected)

	//fmt.Println("---")
	//fmt.Println(actualText)
	//fmt.Println("---")

	if actualText != expectedText {
		actualLines := strings.Split(actualText, "\n")
		expectedLines := strings.Split(expectedText, "\n")

		if len(actualLines) != len(expectedLines) {
			t.Fatalf("count mismatch (actual: %d, expected: %d", len(actualLines), len(expectedLines))
		}

		for line := range expectedLines {
			thisActualLine := strings.TrimSpace(actualLines[line])
			thisExpectedLine := strings.TrimSpace(expectedLines[line])

			matches := thisActualLine == thisExpectedLine
			if !matches {
				szActual := len(thisActualLine)
				szExpected := len(thisExpectedLine)

				if szActual != szExpected {
					t.Logf("  actual:%2d :'%s'", line, thisActualLine)
					t.Logf("expected:%2d :'%s'", line, thisExpectedLine)
					t.Fatalf("%d: line lengths actual (%d) expected (%d)", line, szActual, szExpected)
				}
			}
		}
		t.Fatal("help text does not match expected text")
	}
}
