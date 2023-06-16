package descriptor

import (
	"github.com/sam-caldwell/go/v2/projects/argparse/types"
	"testing"
)

func TestDescriptor_storeRequired(t *testing.T) {
	test := func(typ types.ArgTypes, value any) {
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
		case types.Boolean:
			_, ok = value.(bool)
		case types.Flag:
			_, ok = value.(bool)
		case types.Float:
			_, ok = value.(float64)
		case types.Integer:
			_, ok = value.(int)
		case types.String:
			_, ok = value.(string)
		default:
			t.Fatal("unknown type")
		}
		if !ok {
			t.Fatalf("Typecheck failed after StoreRequired (%s) '%v'", typ.String(), value)
		}
	}

	test(types.Boolean, true)
	test(types.Boolean, false)
	test(types.Flag, true)
	test(types.Flag, false)
	test(types.Float, -1.0)
	test(types.Float, 0.0)
	test(types.Float, 1.0)
	test(types.Integer, -1)
	test(types.Integer, 0)
	test(types.Integer, 1337)
	test(types.String, "")
	test(types.String, "this is a test string")
}
