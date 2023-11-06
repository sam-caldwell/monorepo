package crsce

import "testing"

/*
 * CRSCE Compress Arguments test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Test the Compress program arguments parser.
 */

func TestArguments_Struct(t *testing.T) {
	var arg Arguments
	if arg.In != "" {
		t.Fatal("expect arg.In empty")
	}
	if arg.Out != "" {
		t.Fatal("expect arg.Out empty")
	}
	if arg.BlockSize != 0 {
		t.Fatal("expect arg.BlockSize empty")
	}
}
