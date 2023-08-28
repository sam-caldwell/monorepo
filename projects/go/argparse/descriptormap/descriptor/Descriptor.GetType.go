package descriptor

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
)

// GetType - Return current type
func (arg *Descriptor) GetType() types.ArgTypes {
	return arg.typ
}
