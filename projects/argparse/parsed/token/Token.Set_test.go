package token

import (
	"github.com/sam-caldwell/go/v2/projects/argparse/types"
	"testing"
)

func TestArgumentElement_Set(t *testing.T) {
	var input Token
	if err := input.Set(types.Boolean, true); err != nil {
		t.Fatal(err)
	}
	if err := input.Set(types.Boolean, false); err != nil {
		t.Fatal(err)
	}
	if err := input.Set(types.Boolean, "bad"); err == nil {
		t.Fatal(err)
	}
	if err := input.Set(types.Flag, true); err != nil {
		t.Fatal(err)
	}
	if err := input.Set(types.Flag, false); err != nil {
		t.Fatal(err)
	}
	if err := input.Set(types.Integer, -1); err != nil {
		t.Fatal(err)
	}
	if err := input.Set(types.Integer, 0); err != nil {
		t.Fatal(err)
	}
	if err := input.Set(types.Integer, 1); err != nil {
		t.Fatal(err)
	}
	if err := input.Set(types.Float, -1.0); err != nil {
		t.Fatal(err)
	}
	if err := input.Set(types.Float, 1.0); err != nil {
		t.Fatal(err)
	}
	if err := input.Set(types.Float, 1.0); err != nil {
		t.Fatal(err)
	}
	if err := input.Set(types.String, ""); err != nil {
		t.Fatal(err)
	}
	if err := input.Set(types.String, "test"); err != nil {
		t.Fatal(err)
	}

}
