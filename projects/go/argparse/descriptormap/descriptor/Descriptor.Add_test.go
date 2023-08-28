package descriptor

import (
	types2 "github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
	"github.com/sam-caldwell/monorepo/v2/projects/go/counters"
	"testing"
)

func TestDescriptor_Add(t *testing.T) {

	test := func(pos *counters.Conditional, short, long string, typ types2.ArgTypes, required bool, dValue any, help string) {
		var argDesc Descriptor

		if err := argDesc.Add(pos, short, long, typ, required, dValue, help); err != nil {
			if err != nil {
				t.Fatal(err)
			}
		}
		if (argDesc.pos >= 0) && (argDesc.pos != (pos.Value() - 1)) {
			t.Fatalf("position %d does not match descriptor position %d", pos.Value()-1, argDesc.pos)
		}
	}

	//Expect no issue
	var pos counters.Conditional
	test(&pos, "", "", types2.Boolean, true, true, "test help0")
	test(&pos, "", "", types2.Boolean, true, true, "test help1")
	test(&pos, "", "", types2.Boolean, true, true, "test help2")
	test(&pos, "", "", types2.Boolean, true, true, "test help3")
	test(&pos, "", "", types2.Boolean, true, true, "test help4")

	test(&pos, "-a", "--all", types2.Boolean, true, true, "test help5")

	//Expect error: -a is a duplicate.
	test(&pos, "-a", "--all", types2.Boolean, true, true, "test help6")

}
