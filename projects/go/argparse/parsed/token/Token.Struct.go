package token

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/argparse/types"
)

// Token - this is a parsed descriptor from command-line
type Token struct {
	typ   types.ArgTypes
	value any
}
