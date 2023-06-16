package argparse

import (
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse/descriptormap"
	"github.com/sam-caldwell/go/v2/projects/argparse/argparse/parsed"
	ordered "github.com/sam-caldwell/orderedset/v2"
)

/*
 * This is the main struct used for the argparse package.
 */

// Arguments - Top-Level ArgParse struct
type Arguments struct {
	pos            counters.ConditionalCounter
	descriptors    descriptormap.Map
	programName    string
	preambleText   ordered.Set
	postscriptText ordered.Set
	err            ordered.Set
	results        parsed.Namespace
}
