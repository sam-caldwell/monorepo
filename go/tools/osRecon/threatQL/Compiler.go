package threatQL

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/types"
)

// Compiler - ThreatQL Query compiler.
//
//	This facilitates compiling a human-readable query into a binary query file
//	for use with the client.
type Compiler struct {
	seq      types.Sequence
	queryDir directory.Path
}
