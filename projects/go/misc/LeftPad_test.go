package misc

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/ansi/Tester"
	"testing"
)

func TestLeftPad(t *testing.T) {
	/*
	 * Assumptions:
	 * 1. Given a string, we expect that we can pad to the left of this string a number of white space characters.
	 */
	const (
		space      = ' '
		underscore = '_'
		root       = "test"
	)
	test := ansi.Test(t)
	defer test.Stats()

	var expectedPadding [2]string
	paddingCharacter := [2]rune{space, underscore}

	//Loop through various padding sizes from size of 'root' string to 10x the size of 'root' string
	for expectedSize := len(root) + 1; expectedSize <= 10*len(root); expectedSize++ {
		// each pass where we increment 'i' we will append padding to our expected padding
		// starting from an empty padding string.
		expectedPadding[0] += string(space)
		expectedPadding[1] += string(underscore)

		// test both expected padding 0 and 1 each time.
		for whichPadIndex := 0; whichPadIndex < len(expectedPadding); whichPadIndex++ {
			expected := expectedPadding[whichPadIndex] + root

			output := LeftPad(paddingCharacter[whichPadIndex], root, expectedSize)

			if actualSize := len(output); actualSize != expectedSize {
				test.Fatalf("output length (%d) does not match expected length (%d)", actualSize, expectedSize)
			}

			if output != expected {
				test.Fatalf("output does not match expected output.\n"+
					"  output: '%s'\n"+
					"expected: '%s' (sz:%d)\n",
					output, expected, expectedSize)
			}
		}
	}
	test.Pass("Incrementing padding-size test")

	if output := LeftPad(space, root, len(root)); output != root {
		test.Fatal("same-length test")
	}
	test.Pass("same-length test")

	if output := LeftPad(space, root, 0); output != root {
		test.Fatal("zero-length test")
	}
	test.Pass("zero-length test")

	for i := 0; i < 2*len(root); i++ {
		if output := LeftPad(space, root, len(root)-i); output != root {
			test.Fatal("negative-length test")
		}
	}
	test.Pass("negative-length test")
}
