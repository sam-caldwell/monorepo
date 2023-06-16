package parsed

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/argparse/types"
	"strings"
	"testing"
)

func TestTokenMap_Add(t *testing.T) {
	var set Namespace

	test := func(n string, typ types.ArgTypes, value any, expectRecord bool, typeCheckFail bool) {
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

	test("bool", types.Boolean, true, false, false)
	test("bool", types.Boolean, false, true, false)
	test("bool", types.Integer, 0, true, false)

	test("flag", types.Flag, true, false, false)
	test("flag", types.Flag, false, true, false)
	test("flag", types.Flag, true, true, false)
	test("flag", types.Flag, false, true, false)
	test("flag", types.Flag, -1, true, true)
	test("flag", types.Flag, 0, true, true)
	test("flag", types.Flag, 1, true, true)
	test("flag", types.Flag, false, true, true)
	test("flag", types.Flag, false, true, true)
	test("flag", types.Flag, false, true, true)

	test("float", types.Float, -10.0, false, false)
	test("float", types.Float, 00.00, true, false)
	test("float", types.Float, 10.00, true, false)
	test("float", types.Float, true, true, true)
	test("float", types.Float, false, true, true)
	test("float", types.Float, -1, true, true)
	test("float", types.Float, 0, true, true)
	test("float", types.Float, 1, true, true)
	test("float", types.Float, "bad", true, true)
	test("float", types.Float, "", true, true)
	test("float", types.Float, "-10.0", true, false)
	test("float", types.Float, "-0.00", true, false)
	test("float", types.Float, "10.00", true, false)

	test("integer", types.Integer, -1, false, false)
	test("integer", types.Integer, 0, true, false)
	test("integer", types.Integer, 1, true, false)
	test("integer", types.Integer, true, true, true)
	test("integer", types.Integer, false, true, true)
	test("integer", types.Integer, -1.0, true, true)
	test("integer", types.Integer, 00.0, true, true)
	test("integer", types.Integer, 01.0, true, true)
	test("integer", types.Integer, "bad", true, true)
	test("integer", types.Integer, "", true, true)
	test("integer", types.Integer, "-10.0", true, false)
	test("integer", types.Integer, "-0.00", true, false)
	test("integer", types.Integer, "10.00", true, false)

	test("string", types.String, "ok", false, true)
	test("string", types.String, "", true, true)
	test("string", types.String, -1, true, false)
	test("string", types.String, 0, true, false)
	test("string", types.String, 1, true, false)
	test("string", types.String, true, true, false)
	test("string", types.String, false, true, false)
	test("string", types.String, -1.0, true, false)
	test("string", types.String, 00.0, true, false)
	test("string", types.String, 01.0, true, false)
	test("string", types.String, "-10.0", true, false)
	test("string", types.String, "-0.00", true, false)
	test("string", types.String, "10.00", true, false)
}
