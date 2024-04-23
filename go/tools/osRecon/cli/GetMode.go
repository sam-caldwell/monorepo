package cli

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/types"
	"os"
	"strings"
)

// GetMode - Get the mode (os.Args[1]) and expect client or server
func GetMode() (m types.OperatingMode, err error) {

	var command string

	if len(os.Args) < 2 {
		return m, fmt.Errorf(errors.MissingArguments)
	}

	if command = strings.TrimSpace(os.Args[1]); (command != "client") && (command != "server") {
		return m, fmt.Errorf(errors.InvalidCommand)
	}

	err = m.FromString(command)

	return m, err

}
