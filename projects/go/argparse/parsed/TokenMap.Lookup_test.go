package parsed

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/argparse/parsed/token"
	types2 "github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
	"testing"
)

func TestTokenMap_Lookup(t *testing.T) {
	var set Namespace
	set.data = make(map[string]token.Token)

	test := func(n string, typ types2.ArgTypes, value any, expectRecord bool) {
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
	test("firstArg", types2.Boolean, true, true)
	test("firstArg", types2.Boolean, false, true)
	test("firstArg", types2.Integer, 0, true)
}
