package descriptor

import (
	"github.com/sam-caldwell/go-monorepo/v2/projects/argparse/types"
	"testing"
)

func TestDescriptor_GetType(t *testing.T) {
	var arg Descriptor
	arg.typ = types.String

	if typ := arg.GetType(); typ != types.String {
		t.Fatalf("Descriptor.GetType() failed.  Got: %v", typ)
	}
}
