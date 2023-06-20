package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"runtime"
)

// OpSysVersion - Return the operating system version
func OpSysVersion() (version string, err error) {

	switch runtime.GOOS {

	case words.Darwin:
		version, err = GetVersion()

	case words.Linux:
		version, err = GetVersion()

	case words.Windows:
		version, err = GetVersion()

	default:
		version, err = words.EmptyString, fmt.Errorf(errors.UnsupportedOpsys)
	}

	return version, err
}
