package token

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
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
