package descriptor

import "github.com/sam-caldwell/go/v2/projects/argparse/types"

// GetType - Return current type
func (arg *Descriptor) GetType() types.ArgTypes {
	return arg.typ
}
