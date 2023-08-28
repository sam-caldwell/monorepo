package token

import (
	types2 "github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
	"testing"
)

func TestArgumentElement_SetType(t *testing.T) {
	var input Token
	expectPass := func(typ types2.ArgTypes) {
		if err := input.SetType(typ); err != nil {
			t.Fatal(err)
		}
		if o := input.GetType(); typ != o {
			t.Fatalf("Expected Type not set: %s", o.String())
		}
	}
	expectFail := func(typ types2.ArgTypes) {
		if err := input.SetType(typ); err == nil {
			t.Fatal(err)
		}
	}

	expectPass(types2.Boolean)
	expectPass(types2.Flag)
	expectPass(types2.Float)
	expectPass(types2.Integer)
	expectPass(types2.String)
	expectFail(types2.ArgTypes(99))
}
