package token

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"testing"
)

func TestArgumentElement_SetValue(t *testing.T) {
	var input Token

	test := func(typ types.ArgTypes, value any, expectFail bool) {
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
	test(types.Boolean, true, false)
	test(types.Boolean, false, false)
	test(types.Boolean, "true", true)
	test(types.Boolean, "false", true)
	test(types.Boolean, -1, true)
	test(types.Boolean, 0, true)
	test(types.Boolean, 1, true)
	test(types.Boolean, -1.0, true)
	test(types.Boolean, 0.0, true)
	test(types.Boolean, 1.0, true)
	test(types.Boolean, "bad", true)

	test(types.Flag, true, false)
	test(types.Flag, false, false)
	test(types.Flag, "true", true)
	test(types.Flag, "false", true)
	test(types.Flag, -1, true)
	test(types.Flag, 0, true)
	test(types.Flag, 1, true)
	test(types.Flag, -1.0, true)
	test(types.Flag, 0.0, true)
	test(types.Flag, 1.0, true)
	test(types.Flag, "bad", true)

	test(types.Float, true, true)
	test(types.Float, false, true)
	test(types.Float, "true", true)
	test(types.Float, "false", true)
	test(types.Float, -1, true)
	test(types.Float, 0, true)
	test(types.Float, 1, true)
	test(types.Float, -1.0, false)
	test(types.Float, 0.0, false)
	test(types.Float, 1.0, false)
	test(types.Float, "bad", true)

	test(types.Integer, true, true)
	test(types.Integer, false, true)
	test(types.Integer, "true", true)
	test(types.Integer, "false", true)
	test(types.Integer, -1, false)
	test(types.Integer, 0, false)
	test(types.Integer, 1, false)
	test(types.Integer, -1.0, true)
	test(types.Integer, 0.0, true)
	test(types.Integer, 1.0, true)
	test(types.Integer, "bad", true)

	test(types.String, true, true)
	test(types.String, false, true)
	test(types.String, "true", false)
	test(types.String, "false", false)
	test(types.String, -1, true)
	test(types.String, 0, true)
	test(types.String, 1, true)
	test(types.String, -1.0, true)
	test(types.String, 0.0, true)
	test(types.String, 1.0, true)
	test(types.String, "good", false)
}
