package descriptormap

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse/types"
	"testing"
)

func TestMap_GetByLong(t *testing.T) {
	var m Map

	failIf := func(condition bool, msg string, n ...any) {
		if condition {
			t.Fatalf(msg, n...)
		}
	}

	addRows := func(pos *counters.ConditionalCounter, name, short, long string, typ types.ArgTypes, required bool, dValue any, help string) {
		err := m.Add(pos, name, short, long, typ, required, dValue, help)
		failIf(err != nil, "%v", err)
	}

	verifyRows := func(name, short, long string, typ types.ArgTypes, required bool, dValue any, help string) {
		n, d := m.GetByLong(&long)

		failIf(n == nil, "Error: nil name returned")
		failIf(d == nil, "Error: row not found %s", name)

		failIf(*n != long, "Error long mismatch in response")
		failIf(d.GetShort() != short, "short mismatch: %s (%s):'%s'", name, short, d.GetShort())
		failIf(d.GetLong() != long, "long mismatch: %s (%s):'%s'", name, long, d.GetLong())
		failIf(d.GetType() != typ, "type mismatch: %s (%s):'%s'", name, typ.String(), d.GetType())
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
	var pos counters.ConditionalCounter
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
