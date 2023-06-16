package descriptor

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
)

// Descriptor - A Descriptor struct
type Descriptor struct {
	pos      int            //Expected position (>=0 is a positional arg)
	short    string         // short descriptor
	long     string         // long descriptor
	typ      types.ArgTypes // descriptor type
	required bool           // indicates the descriptor must provide a value
	dValue   any            // provides a default value when not present
	help     string         // provides help text
}
