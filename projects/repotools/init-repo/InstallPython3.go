package repotools

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
	"runtime"
)

func InstallPython3(commandUsage string) string {
	switch runtime.GOOS {
	case words.Darwin:
		return words.EmptyString
	case words.Linux:
		distro, err := systemrecon.GetFamily()
		exit.OnError(err, exit.GeneralError, commandUsage)

		switch distro {

		case words.Ubuntu:
			return "sudo apt-get install python3;sudo apt-get install python3-pip -y"

		case words.Centos:
			exit.OnError(fmt.Errorf(errors.NotImplemented), exit.GeneralError, commandUsage)
			return words.EmptyString

		default:
			exit.OnError(fmt.Errorf(errors.UnsupportedOpsys+errors.Details, distro), exit.GeneralError, commandUsage)
		}

	case words.Windows:
		exit.OnError(fmt.Errorf(errors.NotImplemented), exit.GeneralError, commandUsage)
		return words.EmptyString

	default:
		exit.OnError(fmt.Errorf(errors.UnsupportedOpsys+errors.Details, runtime.GOOS), exit.GeneralError, commandUsage)
	}
	return words.EmptyString
}
