package token

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
)

// GetType - Return the type for the current token
func (arg *Token) GetType() types.ArgTypes {

	return arg.typ

}
