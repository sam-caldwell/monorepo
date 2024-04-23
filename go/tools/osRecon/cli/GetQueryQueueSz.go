package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetQueryQueueSz - Get the value of --queryQueueSz
func GetQueryQueueSz(defaultQueryQueueSz uint16) (uint16, error) {

	queryQueueSz, err := simpleArgs.GetOptionUint16Value("--queryQueueSz", false)

	if err != nil {
		ansi.Red().Printf("Parsing error (%s): %v", "--queryQueueSz", err).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}

	if queryQueueSz == 0 {
		queryQueueSz = defaultQueryQueueSz
	}

	return queryQueueSz, err

}
