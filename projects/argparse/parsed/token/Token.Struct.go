package token

import (
	"github.com/sam-caldwell/argparse/v2/argparse/types"
)

// Token - this is a parsed descriptor from command-line
type Token struct {
	typ   types.ArgTypes
	value any
}
