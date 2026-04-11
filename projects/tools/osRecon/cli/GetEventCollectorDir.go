package cli

import (
	"github.com/sam-caldwell/monorepo/projects/ansi"
	"github.com/sam-caldwell/monorepo/projects/exit"
	"github.com/sam-caldwell/monorepo/projects/fs/directory"
	"github.com/sam-caldwell/monorepo/projects/simpleArgs"
)

// GetEventCollectorDir - get the cli arg for --eventCollectorDir
func GetEventCollectorDir() (result *directory.Path, err error) {

	const argString = "--eventCollectorDir"

	argEventCollectorDir, err := simpleArgs.GetOptionValue(argString)

	if err != nil {
		ansi.Red().
			Printf("Parsing error (%s): %v", argString, err.Error()).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}

	result, err = directory.NewPath(argEventCollectorDir, true)

	return result, err

}
