package runcommand

import (
	"os/exec"
)

type RealCommandExecutor struct{}

func (r RealCommandExecutor) Execute(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	return cmd.Output()
}
