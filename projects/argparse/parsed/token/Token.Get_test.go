package token

import (
	"github.com/sam-caldwell/go-monorepo/v2/projects/argparse/types"
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
