package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetEventQueueSz - Get the value of --eventQueueSz
func GetEventQueueSz(defaultEventQueueSz uint16) (uint16, error) {
	const argString = "--eventQueueSz"

	eventQueueSz, err := simpleArgs.GetOptionUint16Value(argString, false)

	if err != nil {
		ansi.Red().Printf("Parsing error (%s): %v", argString, err).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}

	if eventQueueSz == 0 {
		eventQueueSz = defaultEventQueueSz
	}

	return eventQueueSz, err

}
