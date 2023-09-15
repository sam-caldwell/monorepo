package systemrecon

/*
 * OpSysVersion ()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * OpSysFamily()
 *
 * 	Return the current operating system version.
 */
import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/misc/words"
)

// OpSysVersion - Return the operating system version
func OpSysVersion() (version string, err error) {

	var opsys string

	if opsys, err = OpSys(); err != nil {
		return version, err
	}

	switch opsys {

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
