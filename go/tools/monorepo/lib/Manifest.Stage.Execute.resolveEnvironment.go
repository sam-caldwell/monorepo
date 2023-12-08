package monorepo

import (
	"fmt"
	"os"
	"os/exec"
)

// resolveEnvironment - Given a list of key-value environment variables and command shell, fix up the environment.
func resolveEnvironment(shell *exec.Cmd, envList *[]EnvironmentVariable) {
	for _, env := range *envList {
		shell.Env = append(os.Environ(),
			fmt.Sprintf("%s=%s", env.Key, env.Value))
	}
}
