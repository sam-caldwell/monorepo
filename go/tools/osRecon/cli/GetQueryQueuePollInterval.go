package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetQueryQueuePollInterval - Get the value of --queryQueuePollInterval
func GetQueryQueuePollInterval(defaultQueryQueuePollInterval uint16) (uint16, error) {

	const argString = "--queryQueuePollInterval"

	queryQueuePollInterval, err := simpleArgs.GetOptionUint16Value(argString, false)

	if err != nil {
		ansi.Red().Printf("Parsing error (%s): %v", argString, err).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}

	if queryQueuePollInterval == 0 {
		queryQueuePollInterval = defaultQueryQueuePollInterval
	}

	return queryQueuePollInterval, err

}
