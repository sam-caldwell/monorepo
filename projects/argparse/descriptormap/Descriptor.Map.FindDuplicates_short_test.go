package descriptormap

import (
	"fmt"
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"github.com/sam-caldwell/counters/v2"
	"testing"
)

func TestMap_FindDuplicates_short(t *testing.T) {
	var m Map
	test := func(pos *counters.ConditionalCounter, name, short, long string, typ types.ArgTypes, required bool, dValue any, help string) {
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
	testList := map[string]map[string]map[string]map[any]types.ArgTypes{
		"boolean": {
			"-1": {
				"--all": {
					true:  types.Boolean,
					false: types.Boolean,
				},
			},
		},
		"flag1": {
			"-f": {
				"--fall1": {
					true:  types.Boolean,
					false: types.Boolean,
				},
			},
		},
		"flag2": {
			"-f": {
				"--fall2": {
					true:  types.Boolean,
					false: types.Boolean,
				},
			},
		},
		"float": {
			"-2": {
				"--float": {
					-1.0: types.Float,
					0.0:  types.Float,
					1.0:  types.Float,
				},
			},
		},
		"integer": {
			"-3": {
				"--integer": {
					-1: types.Integer,
					0:  types.Integer,
					1:  types.Integer,
				},
			},
		},
		"string": {
			"-4": {
				"--string": {
					"test1": types.String,
					"test2": types.String,
					"test3": types.String,
				},
			},
		},
	}

	var pos counters.ConditionalCounter
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
