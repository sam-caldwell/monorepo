package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetQueryCollectorDir - get the cli arg --queryCollectorDir
func GetQueryCollectorDir() (result *directory.Path, err error) {

	const argString = "--queryCollectorDir"

	argRaw, err := simpleArgs.GetOptionValue(argString)

	if err != nil {
		ansi.Red().
			Printf("Parsing error (%s): %v", argString, err.Error()).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}

	result, err = directory.NewPath(argRaw, true)

	return result, err

}
