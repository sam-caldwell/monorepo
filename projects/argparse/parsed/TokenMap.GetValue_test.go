package parsed

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"testing"
)

func TestArgumentElementMap_GetValue(t *testing.T) {

	var set Namespace

	//test() is the test function that takes a name, type and value
	test := func(n string, typ types.ArgTypes, value any) {
		//Add element to the Namespace.Add()
		err := set.Add(&n, typ, &value)
		if err != nil {
			t.Fatal(err)
		}
		//Get the added token and verify the result.
		if o, err := set.GetValue(&n); err != nil {
			t.Fatal("expected record not found")
		} else {
			if tcErr := typ.Typecheck(o); tcErr != nil {
				t.Fatal(tcErr)
			}
			if o != value {
				t.Fatal("error: value mismatch")
			}
		}
	}

	test("bool", types.Boolean, true)
	test("bool", types.Boolean, false)

	test("flag", types.Flag, true)
	test("flag", types.Flag, false)
	test("flag", types.Flag, true)
	test("flag", types.Flag, false)

	test("float", types.Float, -10.0)
	test("float", types.Float, 00.00)
	test("float", types.Float, 10.00)

	test("integer", types.Integer, -1)
	test("integer", types.Integer, 0)
	test("integer", types.Integer, 1)

	test("string", types.String, "ok")
	test("string", types.String, "")
	test("string", types.String, "-10.0")
	test("string", types.String, "-0.00")
	test("string", types.String, "10.00")
}
