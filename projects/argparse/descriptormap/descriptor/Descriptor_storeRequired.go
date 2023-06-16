package descriptor

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
)

// storeRequired - store the required flag and type-check the default value
func (arg *Descriptor) storeRequired(t types.ArgTypes, required bool, dValue any) (err error) {
	arg.dValue = dValue
	arg.required = required
	arg.typ = t
	return t.Typecheck(dValue)
}
