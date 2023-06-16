package token

import (
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse/types"
)

// Token - this is a parsed descriptor from command-line
type Token struct {
	typ   types.ArgTypes
	value any
}
