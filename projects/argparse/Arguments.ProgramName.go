package argparse

import (
	"os"
	"path"
)

// ProgramName - Set the program name from os.Args[0]
func (arg *Arguments) ProgramName() *Arguments {
	arg.programName = path.Base(os.Args[0])
	return arg
}
