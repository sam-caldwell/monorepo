package descriptormap

import (
	"fmt"
	types2 "github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
	"github.com/sam-caldwell/monorepo/v2/projects/go/counters"
	"testing"
)

func TestMap_CountPositionalArgs(t *testing.T) {
	var m Map

	addRows := func(pos *counters.Conditional, name, short, long string, typ types2.ArgTypes, required bool, dValue any, help string) {
		err := m.Add(pos, name, short, long, typ, required, dValue, help)
		if err != nil {
			t.Fatal(err)
		}
	}
	validiate := func() {
		if count := m.CountPositionalArgs(); count != 3 {
			t.Fatal("Count wrong")
		}
	}

	//map[name]map[short]map[long]map[dValue]typ
	testList := map[string]map[string]map[string]map[any]types2.ArgTypes{
		"pos0": {
			"": {
				"": {
					true:  types2.Boolean,
					false: types2.Boolean,
				},
			},
		},
		"pos1": {
			"": {
				"": {
					true:  types2.Boolean,
					false: types2.Boolean,
				},
			},
		},
		"pos2": {
			"": {
				"": {
					true:  types2.Boolean,
					false: types2.Boolean,
				},
			},
		},
		"boolean": {
			"-a": {
				"--all": {
					true:  types2.Boolean,
					false: types2.Boolean,
				},
			},
		},
		"flag": {
			"-f": {
				"--fall": {
					true:  types2.Boolean,
					false: types2.Boolean,
				},
			},
		},
		"float": {
			"-n": {
				"--float": {
					-1.0: types2.Float,
					0.0:  types2.Float,
					1.0:  types2.Float,
				},
			},
		},
		"integer": {
			"-i": {
				"--integer": {
					-1: types2.Integer,
					0:  types2.Integer,
					1:  types2.Integer,
				},
			},
		},
		"string": {
			"-s": {
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
						addRows(&pos, name, short, long, typ, required, value, fmt.Sprintf("help text %s", name))
					}
				}
			}
		}
	}
	validiate()
}
