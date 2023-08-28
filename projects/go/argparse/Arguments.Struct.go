package argparse

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/argparse/descriptormap"
	"github.com/sam-caldwell/monorepo/v2/projects/go/argparse/parsed"
	"github.com/sam-caldwell/monorepo/v2/projects/go/counters"
	"github.com/sam-caldwell/monorepo/v2/projects/go/sets/ordered"
)

/*
 * This is the main struct used for the argparse package.
 */

// Arguments - Top-Level ArgParse struct
type Arguments struct {
	pos         counters.Conditional
	descriptors descriptormap.Map
	programName    string
	preambleText   ordered.Set
	postscriptText ordered.Set
	err            ordered.Set
	results        parsed.Namespace
}
