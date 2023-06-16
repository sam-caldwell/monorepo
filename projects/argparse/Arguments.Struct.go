package argparse

import (
	"github.com/sam-caldwell/go/v2/projects/argparse/descriptormap"
	"github.com/sam-caldwell/go/v2/projects/argparse/parsed"
	"github.com/sam-caldwell/go/v2/projects/counters"
	"github.com/sam-caldwell/go/v2/projects/sets/ordered"
)

/*
 * This is the main struct used for the argparse package.
 */

// Arguments - Top-Level ArgParse struct
type Arguments struct {
	pos            counters.Conditional
	descriptors    descriptormap.Map
	programName    string
	preambleText   ordered.Set
	postscriptText ordered.Set
	err            ordered.Set
	results        parsed.Namespace
}
