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
		distro, err := systemrecon.GetVersion()
		exit.OnError(err, exit.GeneralError, commandUsage)

		switch distro {

		case words.Ubuntu:
			return "command -v python3 || sudo apt-get install python3 python3-pip -y"

		case words.Centos:
			exit.OnError(fmt.Errorf(errors.NotImplemented), exit.GeneralError, commandUsage)
			return words.EmptyString

		default:
			exit.OnError(fmt.Errorf(errors.UnsupportedOpsys), exit.GeneralError, commandUsage)
		}

	case words.Windows:
		exit.OnError(fmt.Errorf(errors.NotImplemented), exit.GeneralError, commandUsage)
		return words.EmptyString

	default:
		exit.OnError(fmt.Errorf(errors.UnsupportedOpsys), exit.GeneralError, commandUsage)
	}
	return words.EmptyString
}
