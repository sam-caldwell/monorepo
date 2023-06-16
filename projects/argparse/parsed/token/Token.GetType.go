package token

import (
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse/types"
)

// GetType - Return the type for the current token
func (arg *Token) GetType() types.ArgTypes {

	return arg.typ

}
