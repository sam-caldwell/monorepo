package token

import (
	"github.com/sam-caldwell/go-monorepo/v2/projects/argparse/types"
)

// GetType - Return the type for the current token
func (arg *Token) GetType() types.ArgTypes {

	return arg.typ

}
