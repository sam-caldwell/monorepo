package descriptormap

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
	"github.com/sam-caldwell/counters/v2"
	"testing"
)

func TestMap_GetByPosition(t *testing.T) {
	var m Map
	var pos counters.ConditionalCounter

	addRow := func(pos *counters.ConditionalCounter, name, short, long string, typ types.ArgTypes, required bool, dValue any, help string) {
		if err := m.Add(pos, name, short, long, typ, required, dValue, help); err != nil {
			t.Fatalf("%v", err)
		}
	}

	validate := func() {
		if m.data == nil {
			t.Fatal("Setup is wrong")
		}
		for expectedName, expectedArg := range m.data {
			//Lookup position for current descriptor
			actualName, actualArg := m.GetByPosition(expectedArg.GetPosition())
			if actualName == nil {
				t.Fatalf("%s actualName at position %d not found", expectedName, expectedArg.GetPosition())
			}
			if actualArg == nil {
				t.Fatalf("%s actualArg at position %d not found", expectedName, expectedArg.GetPosition())
			}
			if actualArg.GetPosition() != expectedArg.GetPosition() {
				t.Fatalf("%s actual and expected position mismatch", expectedName)
			}
			switch expectedArg.GetHelp() {
			case "positional":
				if actualArg.GetPosition() < 0 {
					t.Fatalf("%s positional args should have position >=0", expectedName)
				}
			case "optional":
				if actualArg.GetPosition() != -1 {
					t.Fatalf("%s optional arguments should return position -1", expectedName)
				}
			default:
				t.Fatalf("%s test is not setup right (help)", expectedName)
			}
		}
		for i := 0; i < 3; i++ {
			actualName, actualArg := m.GetByPosition(i)
			if actualName == nil {
				t.Fatalf("actualName at position %d not found", i)
			}
			if actualArg == nil {
				t.Fatalf("actualArg at position %d not found", i)
			}
			if i != actualArg.GetPosition() {
				t.Fatalf("actualArg position (%d) does not match expected position (%d)",
					actualArg.GetPosition(), i)
			}

		}
	}

	addRow(&pos, "pos0", "", "", types.Boolean, true, false, "positional")
	addRow(&pos, "pos1", "", "", types.Boolean, true, false, "positional")
	addRow(&pos, "pos2", "", "", types.Boolean, true, false, "positional")
	addRow(&pos, "opt0", "", "", types.Boolean, true, false, "positional")
	addRow(&pos, "opt4", "-4", "--opt4", types.Integer, true, 4, "optional")
	addRow(&pos, "opt5", "-5", "--opt5", types.Integer, true, 5, "optional")
	addRow(&pos, "opt6", "-6", "--opt6", types.Integer, true, 6, "optional")
	addRow(&pos, "opt7", "-7", "--opt7", types.Integer, true, 7, "optional")
	addRow(&pos, "opt8", "-8", "--opt8", types.Integer, true, 8, "optional")
	validate()
}
