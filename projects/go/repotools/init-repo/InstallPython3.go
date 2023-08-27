package repotools

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	"github.com/sam-caldwell/go/v2/projects/go/exit/errors"
	words2 "github.com/sam-caldwell/go/v2/projects/go/misc/words"
	"github.com/sam-caldwell/go/v2/projects/go/systemrecon/opsys"
	"runtime"
)

func InstallPython3(commandUsage string) string {
	switch runtime.GOOS {
	case words2.Darwin:
		return words2.EmptyString
	case words2.Linux:
		distro, err := systemrecon.GetFamily()
		exit.OnError(err, exit.GeneralError, commandUsage)

		switch distro {

		case words2.Ubuntu:
			return "sudo apt-get install python3;sudo apt-get install python3-pip -y"

		case words2.Centos:
			exit.OnError(fmt.Errorf(errors.NotImplemented), exit.GeneralError, commandUsage)
			return words2.EmptyString

		default:
			exit.OnError(fmt.Errorf(errors.UnsupportedOpsys+errors.Details, distro), exit.GeneralError, commandUsage)
		}

	case words2.Windows:
		exit.OnError(fmt.Errorf(errors.NotImplemented), exit.GeneralError, commandUsage)
		return words2.EmptyString

	default:
		exit.OnError(fmt.Errorf(errors.UnsupportedOpsys+errors.Details, runtime.GOOS), exit.GeneralError, commandUsage)
	}
	return words2.EmptyString
}
