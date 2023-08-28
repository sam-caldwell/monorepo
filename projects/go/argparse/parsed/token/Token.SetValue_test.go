package token

import (
	types2 "github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
	"testing"
)

func TestArgumentElement_SetValue(t *testing.T) {
	var input Token

	test := func(typ types2.ArgTypes, value any, expectFail bool) {
		if err := input.SetType(typ); err != nil {
			t.Fatal(err)
		}
		if err := input.SetValue(value); err != nil {
			if !expectFail {
				t.Fatal(err)
			}
		}
		if actualValue := input.GetValue(); actualValue != value {
			if !expectFail {
				t.Fatal("Stored value does not match input")
			}
		}
	}
	test(types2.Boolean, true, false)
	test(types2.Boolean, false, false)
	test(types2.Boolean, "true", true)
	test(types2.Boolean, "false", true)
	test(types2.Boolean, -1, true)
	test(types2.Boolean, 0, true)
	test(types2.Boolean, 1, true)
	test(types2.Boolean, -1.0, true)
	test(types2.Boolean, 0.0, true)
	test(types2.Boolean, 1.0, true)
	test(types2.Boolean, "bad", true)

	test(types2.Flag, true, false)
	test(types2.Flag, false, false)
	test(types2.Flag, "true", true)
	test(types2.Flag, "false", true)
	test(types2.Flag, -1, true)
	test(types2.Flag, 0, true)
	test(types2.Flag, 1, true)
	test(types2.Flag, -1.0, true)
	test(types2.Flag, 0.0, true)
	test(types2.Flag, 1.0, true)
	test(types2.Flag, "bad", true)

	test(types2.Float, true, true)
	test(types2.Float, false, true)
	test(types2.Float, "true", true)
	test(types2.Float, "false", true)
	test(types2.Float, -1, true)
	test(types2.Float, 0, true)
	test(types2.Float, 1, true)
	test(types2.Float, -1.0, false)
	test(types2.Float, 0.0, false)
	test(types2.Float, 1.0, false)
	test(types2.Float, "bad", true)

	test(types2.Integer, true, true)
	test(types2.Integer, false, true)
	test(types2.Integer, "true", true)
	test(types2.Integer, "false", true)
	test(types2.Integer, -1, false)
	test(types2.Integer, 0, false)
	test(types2.Integer, 1, false)
	test(types2.Integer, -1.0, true)
	test(types2.Integer, 0.0, true)
	test(types2.Integer, 1.0, true)
	test(types2.Integer, "bad", true)

	test(types2.String, true, true)
	test(types2.String, false, true)
	test(types2.String, "true", false)
	test(types2.String, "false", false)
	test(types2.String, -1, true)
	test(types2.String, 0, true)
	test(types2.String, 1, true)
	test(types2.String, -1.0, true)
	test(types2.String, 0.0, true)
	test(types2.String, 1.0, true)
	test(types2.String, "good", false)
}
