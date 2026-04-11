package cli

import (
	"github.com/sam-caldwell/monorepo/projects/ansi"
	"github.com/sam-caldwell/monorepo/projects/exit"
	"github.com/sam-caldwell/monorepo/projects/simpleArgs"
)

// GetQueryQueueSz - Get the value of --queryQueueSz
func GetQueryQueueSz(defaultQueryQueueSz uint16) (uint16, error) {

	const argString = "--queryQueueSz"
	queryQueueSz, err := simpleArgs.GetOptionUint16Value(argString, false)

	if err != nil {
		ansi.Red().Printf("Parsing error (%s): %v", argString, err).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}

	if queryQueueSz == 0 {
		queryQueueSz = defaultQueryQueueSz
	}

	return queryQueueSz, err

}
