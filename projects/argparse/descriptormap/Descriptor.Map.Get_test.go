package descriptormap

import (
	"fmt"
	"github.com/sam-caldwell/go-monorepo/v2/projects/argparse/types"
	"github.com/sam-caldwell/go-monorepo/v2/projects/counters"
	"testing"
)

func TestMap_Get(t *testing.T) {
	var m Map

	addRows := func(pos *counters.Conditional, name, short, long string, typ types.ArgTypes, required bool, dValue any, help string) {
		err := m.Add(pos, name, short, long, typ, required, dValue, help)
		if err != nil {
			t.Fatal(err)
		}
	}
	verifyRows := func(name, short, long string, typ types.ArgTypes, required bool, dValue any, help string) {
		var d = m.Get(&name)
		if d == nil {
			t.Fatalf("Error: row not found %s", name)
		} else {
			if d.GetShort() != short {
				t.Fatalf("short mismatch:%s", d.GetShort())
			}
			if d.GetLong() != long {
				t.Fatalf("long mismatch:%s", d.GetLong())
			}
		}
	}

	//map[name]map[short]map[long]map[dValue]typ
	testList := map[string]map[string]map[string]map[any]types.ArgTypes{
		"boolean": {
			"-a": {
				"--all": {
					true:  types.Boolean,
					false: types.Boolean,
				},
			},
		},
		"flag": {
			"-f": {
				"--fall": {
					true:  types.Boolean,
					false: types.Boolean,
				},
			},
		},
		"float": {
			"-n": {
				"--float": {
					-1.0: types.Float,
					0.0:  types.Float,
					1.0:  types.Float,
				},
			},
		},
		"integer": {
			"-i": {
				"--integer": {
					-1: types.Integer,
					0:  types.Integer,
					1:  types.Integer,
				},
			},
		},
		"string": {
			"-s": {
				"--string": {
					"test1": types.String,
					"test2": types.String,
					"test3": types.String,
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
	for name, level1 := range testList {
		for short, level2 := range level1 {
			for long, level3 := range level2 {
				for value, typ := range level3 {
					for _, required := range []bool{true, false} {
						verifyRows(name, short, long, typ, required, value, fmt.Sprintf("help text %s", name))
					}
				}
			}
		}
	}
}
