package parsed

import (
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse/parsed/token"
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse/types"
	"testing"
)

func TestTokenMap_Lookup(t *testing.T) {
	var set Namespace
	set.data = make(map[string]token.Token)

	test := func(n string, typ types.ArgTypes, value any, expectRecord bool) {
		var element token.Token
		if err := element.Set(typ, value); err != nil {
			t.Fatal(err)
		}
		set.data[n] = element
		searchResult := set.Lookup(&n)

		if searchResult == nil {
			t.Fatalf("expected record but encountered none (%s)", n)
		} else {
			return
		}
	}
	test("firstArg", types.Boolean, true, true)
	test("firstArg", types.Boolean, false, true)
	test("firstArg", types.Integer, 0, true)
}
