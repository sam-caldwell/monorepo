package argparse

import (
	"github.com/sam-caldwell/go-monorepo/v2/projects/argparse/parsed"
)

// Reduce - Reduce the parsed arguments to a Namespace struct and return the same
func (arg *Arguments) Reduce() (result parsed.Namespace) {
	return arg.results
}
