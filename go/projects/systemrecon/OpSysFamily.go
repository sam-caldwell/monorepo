package systemrecon

/*
 * OpSysFamily ()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * OpSysFamily()
 *
 * 	Return the current operating system family (based on GOOS opsys detection)
 */
import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/exit/errors"
	"github.com/sam-caldwell/monorepo/go/projects/misc/words"
)

// OpSysFamily - return the operating system family (e.g. Ubuntu, Redhat, Centos)
func OpSysFamily() (family string, err error) {

	var opsys string

	if opsys, err = OpSys(); err != nil {
		return family, err
	}

	switch opsys {

	case words.Linux:
		return GetFamily()

	case words.Windows:
		return GetFamily()

	case words.Darwin:
		return GetFamily()

	default:
		return opsys, fmt.Errorf(errors.UnsupportedOpsys)

	}

}
