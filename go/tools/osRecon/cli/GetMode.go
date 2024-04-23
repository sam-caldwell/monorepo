package cli

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/types"
	"os"
)

// GetMode - Get the mode (os.Args[1]) and expect client or server
func GetMode() (m types.OperatingMode, err error) {

	var command string

	if len(os.Args) < 2 {
		return m, fmt.Errorf(errors.MissingArguments)
	}

	err = m.FromString(command)

	return m, err

}
