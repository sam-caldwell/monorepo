package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go-monorepo/v2/projects/argparse/types"
	"testing"
)

func TestArguments_Add(t *testing.T) {
	var arg Arguments

	if count := arg.descriptors.Count(); count != 0 {
		t.Fatalf("Expected 0 records. Got %d", count)
	}

	p := 0
	for i := 0; i < 4; i++ {
		for n, required := range []bool{true, false} {
			name := fmt.Sprintf("name%d%d", i, n)
			short := fmt.Sprintf("-%d", p)
			long := fmt.Sprintf("--arg%d%d", i, n)
			typ := types.Boolean
			value := required
			help := fmt.Sprintf("help string %d%d", i, n)
			arg.Add(name, short, long, typ, required, value, help)
			p++
		}
	}

	if count := arg.descriptors.Count(); count != p {
		t.Fatalf("Expected %d records. Got %d", p, count)
	}

	if count := arg.ErrorCount(); count != 0 {
		t.Log("error count has errors")
		for i, e := range arg.Errors() {
			t.Logf("error(%d):%s", i, e)
		}
		t.Fatalf("Expected no errors.  Got: %d", count)
	}

	if arg.Get("name00").GetShort() != "-0" {
		t.Fatal("expected name0.0 to have -0 short arg")
	}

	p = 0
	for i := 0; i < 4; i++ {
		for n := range []bool{true, false} {
			name := fmt.Sprintf("name%d%d", i, n)
			short := fmt.Sprintf("-%d", p)
			long := fmt.Sprintf("--arg%d%d", i, n)
			help := fmt.Sprintf("help string %d%d", i, n)
			if arg.Get(name).GetShort() != short {
				t.Fatalf("expected %s to have %s short arg", name, short)
			}
			if arg.Get(name).GetLong() != long {
				t.Fatalf("expected %s to have %s long arg", name, long)
			}
			if arg.Get(name).GetHelp() != help {
				t.Fatalf("expected %s to have %s help string", name, help)
			}
			p++
		}
	}
}
