package descriptor

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
	"testing"
)

func TestDescriptor_storeType(t *testing.T) {
	var descriptor Descriptor

	if err := descriptor.storeType(types.Boolean); err != nil {
		t.Fatal(err)
	}
	if descriptor.typ != types.Boolean {
		t.Fatal(descriptor.typ.String())
	}

	if err := descriptor.storeType(types.Flag); err != nil {
		t.Fatal(err)
	}
	if descriptor.typ != types.Flag {
		t.Fatal(descriptor.typ.String())
	}

	if err := descriptor.storeType(types.Float); err != nil {
		t.Fatal(err)
	}
	if descriptor.typ != types.Float {
		t.Fatal(descriptor.typ.String())
	}

	if err := descriptor.storeType(types.Integer); err != nil {
		t.Fatal(err)
	}
	if descriptor.typ != types.Integer {
		t.Fatal(descriptor.typ.String())
	}

	if err := descriptor.storeType(types.String); err != nil {
		t.Fatal(err)
	}
	if descriptor.typ != types.String {
		t.Fatal(descriptor.typ.String())
	}

}
