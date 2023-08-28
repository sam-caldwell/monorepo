package token

import (
	types2 "github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
	"testing"
)

func TestArgumentElement_GetType(t *testing.T) {
	var input Token
	data := []types2.ArgTypes{
		types2.Boolean,
		types2.String,
		types2.Integer,
		types2.Flag,
		types2.Float,
		types2.String,
	}
	for _, thisT := range data {
		input.typ = thisT
		if input.typ != thisT {
			t.Fatalf("GetType mismatch (%s)(%s)",
				input.typ.String(), thisT.String())
		}
	}
}
