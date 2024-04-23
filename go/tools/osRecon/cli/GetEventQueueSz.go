package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetEventQueueSz - Get the value of --eventQueueSz
func GetEventQueueSz(defaultEventQueueSz uint16) (uint16, error) {

	eventQueueSz, err := simpleArgs.GetOptionUint16Value("--eventQueueSz", false)

	if err != nil {
		ansi.Red().Printf("Parsing error (%s): %v", "--eventQueueSz", err).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}

	if eventQueueSz == 0 {
		eventQueueSz = defaultEventQueueSz
	}

	return eventQueueSz, err

}
