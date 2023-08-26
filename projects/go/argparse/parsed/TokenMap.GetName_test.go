package parsed

import (
	types2 "github.com/sam-caldwell/go/v2/projects/go/argparse/types"
	"testing"
)

func TestArgumentElementMap_GetName(t *testing.T) {

	var set Namespace

	//test() is the test function that takes a name, type and value
	test := func(n string, typ types2.ArgTypes, value any) {
		//Add element to the Namespace.Add()
		err := set.Add(&n, typ, &value)
		if err != nil {
			t.Fatal(err)
		}
		//Get the added token and verify the result.
		if o, err := set.GetName(&n); (err != nil) || (o == "") {
			t.Fatal("expected record not found")
		}
	}

	test("bool", types2.Boolean, true)
	test("bool", types2.Boolean, false)

	test("flag", types2.Flag, true)
	test("flag", types2.Flag, false)
	test("flag", types2.Flag, true)
	test("flag", types2.Flag, false)

	test("float", types2.Float, -10.0)
	test("float", types2.Float, 00.00)
	test("float", types2.Float, 10.00)

	test("integer", types2.Integer, -1)
	test("integer", types2.Integer, 0)
	test("integer", types2.Integer, 1)

	test("string", types2.String, "ok")
	test("string", types2.String, "")
	test("string", types2.String, "-10.0")
	test("string", types2.String, "-0.00")
	test("string", types2.String, "10.00")
}
