package monorepo

import (
	"fmt"
	"os"
	"os/exec"
)

// setGolangEnvParams - If the className is 'go' set the shell environment variables
func setGolangEnvParams(shell *exec.Cmd, className, opsys, arch *string) {
	//Go-specific environment setup
	if *className == "go" {
		shell.Env = append(os.Environ(),
			fmt.Sprintf("GOOS=%s", *opsys),
			fmt.Sprintf("GOARCH=%s", *arch),
		)
	}
}
