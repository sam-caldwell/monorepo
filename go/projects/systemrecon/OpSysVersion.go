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
	"github.com/sam-caldwell/monorepo/go/projects/v2/exit/errors"
	words2 "github.com/sam-caldwell/monorepo/go/projects/v2/misc/words"
)

// OpSysVersion - Return the operating system version
func OpSysVersion() (version string, err error) {

	var opsys string

	if opsys, err = OpSys(); err != nil {
		return version, err
	}

	switch opsys {

	case words2.Darwin:
		version, err = GetVersion()

	case words2.Linux:
		version, err = GetVersion()

	case words2.Windows:
		version, err = GetVersion()

	default:
		version, err = words2.EmptyString, fmt.Errorf(errors.UnsupportedOpsys)
	}

	return version, err
}
