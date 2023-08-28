package descriptormap

import (
	"fmt"
	types2 "github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
	"github.com/sam-caldwell/monorepo/v2/projects/go/counters"
	"testing"
)

func TestMap_FindDuplicates_short(t *testing.T) {
	var m Map
	test := func(pos *counters.Conditional, name, short, long string, typ types2.ArgTypes, required bool, dValue any, help string) {
		err := m.Add(pos, name, short, long, typ, required, dValue, help)
		if err != nil {
			t.Fatal(err)
		}
	}
	validate := func() {
		if err := m.FindDuplicates(); err != nil {
			if err.Error() != fmt.Sprintf(errArgumentCannotBeRedefined, "flag2") {
				t.Fatalf("Expected to find duplicate short")
			}
		}
	}

	//map[name]map[short]map[long]map[dValue]typ
	testList := map[string]map[string]map[string]map[any]types2.ArgTypes{
		"boolean": {
			"-1": {
				"--all": {
					true:  types2.Boolean,
					false: types2.Boolean,
				},
			},
		},
		"flag1": {
			"-f": {
				"--fall1": {
					true:  types2.Boolean,
					false: types2.Boolean,
				},
			},
		},
		"flag2": {
			"-f": {
				"--fall2": {
					true:  types2.Boolean,
					false: types2.Boolean,
				},
			},
		},
		"float": {
			"-2": {
				"--float": {
					-1.0: types2.Float,
					0.0:  types2.Float,
					1.0:  types2.Float,
				},
			},
		},
		"integer": {
			"-3": {
				"--integer": {
					-1: types2.Integer,
					0:  types2.Integer,
					1:  types2.Integer,
				},
			},
		},
		"string": {
			"-4": {
				"--string": {
					"test1": types2.String,
					"test2": types2.String,
					"test3": types2.String,
				},
			},
		},
	}

	var pos counters.Conditional
	for name, level1 := range testList {
		for short, level2 := range level1 {
			for long, level3 := range level2 {
				for value, typ := range level3 {
					for _, required := range []bool{true, false} {
						test(&pos, name, short, long, typ, required, value, fmt.Sprintf("help text %s", name))
					}
				}
			}
		}
	}
	validate()
}
