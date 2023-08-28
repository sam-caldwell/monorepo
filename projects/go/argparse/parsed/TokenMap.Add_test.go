package parsed

import (
	"fmt"
	types2 "github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
	"strings"
	"testing"
)

func TestTokenMap_Add(t *testing.T) {
	var set Namespace

	test := func(n string, typ types2.ArgTypes, value any, expectRecord bool, typeCheckFail bool) {
		if err := set.Add(&n, typ, &value); err != nil {
			actualError := strings.ToLower(err.Error())
			if expectRecord {
				if typeCheckFail {
					element := set.data[n]
					thisType := element.GetType()
					expectedError := strings.ToLower(
						fmt.Errorf("type-check failed (%s)",
							thisType.String()).Error())

					if actualError != expectedError {
						//t.Logf("typeCheckFail:%v", typeCheckFail)
						//t.Logf("       actual:%s", actualError)
						//t.Logf("     expected:%s", expectedError)
						t.Fatal(err)
					}
				} else {
					expectedError := fmt.Errorf(errTypeCannotChange, n)
					//t.Logf("typeCheckFail:%v", typeCheckFail)
					//t.Logf("       actual:%s", actualError)
					//t.Logf("     expected:%s", expectedError)
					if actualError != expectedError.Error() {
						t.Fatal(err)
					}
				}
			} else {
				t.Fatalf("Failed to Add() new element to set. %v", err)
			}
		}
		if expectRecord {
			if set.data != nil {
				if _, found := set.data[n]; !found {
					t.Fatal("Expected record, not found.")
				}
			}
		}
	}

	test("bool", types2.Boolean, true, false, false)
	test("bool", types2.Boolean, false, true, false)
	test("bool", types2.Integer, 0, true, false)

	test("flag", types2.Flag, true, false, false)
	test("flag", types2.Flag, false, true, false)
	test("flag", types2.Flag, true, true, false)
	test("flag", types2.Flag, false, true, false)
	test("flag", types2.Flag, -1, true, true)
	test("flag", types2.Flag, 0, true, true)
	test("flag", types2.Flag, 1, true, true)
	test("flag", types2.Flag, false, true, true)
	test("flag", types2.Flag, false, true, true)
	test("flag", types2.Flag, false, true, true)

	test("float", types2.Float, -10.0, false, false)
	test("float", types2.Float, 00.00, true, false)
	test("float", types2.Float, 10.00, true, false)
	test("float", types2.Float, true, true, true)
	test("float", types2.Float, false, true, true)
	test("float", types2.Float, -1, true, true)
	test("float", types2.Float, 0, true, true)
	test("float", types2.Float, 1, true, true)
	test("float", types2.Float, "bad", true, true)
	test("float", types2.Float, "", true, true)
	test("float", types2.Float, "-10.0", true, false)
	test("float", types2.Float, "-0.00", true, false)
	test("float", types2.Float, "10.00", true, false)

	test("integer", types2.Integer, -1, false, false)
	test("integer", types2.Integer, 0, true, false)
	test("integer", types2.Integer, 1, true, false)
	test("integer", types2.Integer, true, true, true)
	test("integer", types2.Integer, false, true, true)
	test("integer", types2.Integer, -1.0, true, true)
	test("integer", types2.Integer, 00.0, true, true)
	test("integer", types2.Integer, 01.0, true, true)
	test("integer", types2.Integer, "bad", true, true)
	test("integer", types2.Integer, "", true, true)
	test("integer", types2.Integer, "-10.0", true, false)
	test("integer", types2.Integer, "-0.00", true, false)
	test("integer", types2.Integer, "10.00", true, false)

	test("string", types2.String, "ok", false, true)
	test("string", types2.String, "", true, true)
	test("string", types2.String, -1, true, false)
	test("string", types2.String, 0, true, false)
	test("string", types2.String, 1, true, false)
	test("string", types2.String, true, true, false)
	test("string", types2.String, false, true, false)
	test("string", types2.String, -1.0, true, false)
	test("string", types2.String, 00.0, true, false)
	test("string", types2.String, 01.0, true, false)
	test("string", types2.String, "-10.0", true, false)
	test("string", types2.String, "-0.00", true, false)
	test("string", types2.String, "10.00", true, false)
}
