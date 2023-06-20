package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"runtime"
)

// OpSysFamily - return the operating system family (e.g. Ubuntu, Redhat, Centos)
func OpSysFamily() (family string, err error) {
	switch opsys := runtime.GOOS; opsys {
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
