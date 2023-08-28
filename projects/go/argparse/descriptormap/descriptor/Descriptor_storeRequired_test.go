package descriptor

import (
	types2 "github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
	"testing"
)

func TestDescriptor_storeRequired(t *testing.T) {
	test := func(typ types2.ArgTypes, value any) {
		var descriptor Descriptor
		if err := descriptor.storeRequired(typ, true, value); err != nil {
			t.Fatal(err)
		}
		if !descriptor.required {
			t.Fail()
		}
		if descriptor.typ != typ {
			t.Fail()
		}
		if descriptor.dValue != value {
			t.Fail()
		}
		var ok bool
		switch descriptor.typ {
		case types2.Boolean:
			_, ok = value.(bool)
		case types2.Flag:
			_, ok = value.(bool)
		case types2.Float:
			_, ok = value.(float64)
		case types2.Integer:
			_, ok = value.(int)
		case types2.String:
			_, ok = value.(string)
		default:
			t.Fatal("unknown type")
		}
		if !ok {
			t.Fatalf("Typecheck failed after StoreRequired (%s) '%v'", typ.String(), value)
		}
	}

	test(types2.Boolean, true)
	test(types2.Boolean, false)
	test(types2.Flag, true)
	test(types2.Flag, false)
	test(types2.Float, -1.0)
	test(types2.Float, 0.0)
	test(types2.Float, 1.0)
	test(types2.Integer, -1)
	test(types2.Integer, 0)
	test(types2.Integer, 1337)
	test(types2.String, "")
	test(types2.String, "this is a test string")
}
