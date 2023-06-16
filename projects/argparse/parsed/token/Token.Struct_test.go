package token

import (
	"github.com/sam-caldwell/go-monorepo/v2/projects/argparse/types"
	"testing"
)

func TestToken_Struct(t *testing.T) {
	var arg Token
	if arg.typ != types.String {
		t.Fail()
	}
	if arg.value != nil {
		t.Fail()
	}
}
