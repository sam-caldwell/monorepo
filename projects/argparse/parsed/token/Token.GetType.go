package token

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
)

// GetType - Return the type for the current token
func (arg *Token) GetType() types.ArgTypes {

	return arg.typ

}
