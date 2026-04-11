package runcommand

import (
	"os/exec"
)

// Execute - Execute a command and return the command's output/error state
func (r RealCommandExecutor) Execute(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	return cmd.Output()
}
