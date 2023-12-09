package monorepo

import (
	"fmt"
	"os"
)

// resolveEnvironment - Given a list of key-value environment variables and command shell, fix up the environment.
func resolveEnvironment(envList *[]EnvironmentVariable) error {
	if len(*envList) == 0 {
		return nil
	}
	for _, env := range *envList {
		if err := os.Setenv(env.Key, env.Value); err != nil {
			return fmt.Errorf("error setting %s (value %s): %v", env.Key, env.Value, err)
		}
	}
	return nil
}
