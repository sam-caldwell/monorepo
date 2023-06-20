package systemrecon

import (
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"runtime"
)

// OpSysVersion - Return the operating system version
func OpSysVersion() (string, error) {
	switch runtime.GOOS {
	case words.Darwin:
		return GetVersion()
	case words.Linux:
		return GetVersion()
	case words.Windows:
		return GetVersion()

	}
	return errors.UnsupportedOpsys, nil
}
