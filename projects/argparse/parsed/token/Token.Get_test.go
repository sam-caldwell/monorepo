package token

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"testing"
)

func TestArgumentElement_Get(t *testing.T) {
	var input Token
	input.typ = types.Boolean
	input.value = true

	result := input.Get()

	if &input == &result {
		t.Fatal("expected a copy of the Token")
	}
}
