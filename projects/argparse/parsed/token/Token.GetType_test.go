package token

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"testing"
)

func TestArgumentElement_GetType(t *testing.T) {
	var input Token
	data := []types.ArgTypes{
		types.Boolean,
		types.String,
		types.Integer,
		types.Flag,
		types.Float,
		types.String,
	}
	for _, thisT := range data {
		input.typ = thisT
		if input.typ != thisT {
			t.Fatalf("GetType mismatch (%s)(%s)",
				input.typ.String(), thisT.String())
		}
	}
}
