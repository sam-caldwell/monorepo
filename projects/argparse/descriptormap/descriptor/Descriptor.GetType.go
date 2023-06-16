package descriptor

import "github.com/sam-caldwell/argparse/v2/argparse/types"

// GetType - Return current type
func (arg *Descriptor) GetType() types.ArgTypes {
	return arg.typ
}
