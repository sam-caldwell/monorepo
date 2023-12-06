package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"os/exec"
	"strings"
)

func (s *Stage) Execute(debug bool) error {
	for _, step := range s.Steps {
		var args []string
		ansi.White().Printf("  step: %s", step.Command).LF().Reset()
		if strings.TrimSpace(step.Command) == "" {
			continue
		}
		parts := strings.Split(step.Command, " ")
		command := parts[0]
		if len(parts[1:]) > 0 {
			args = parts[1:]
		}
		cmd := exec.Command(command, args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		hasError := strings.Contains(strings.ToLower(string(output)), "error")
		if hasError {
			if step.ContinueOnError {
				if debug {
					ansi.White().
						Print("\tOutput: ").
						Yellow().
						Printf("%s\n", strings.TrimSuffix(string(output), "\n")).
						Printf("\tContinueOnError: %v\n", step.ContinueOnError).
						Reset()
				}
			} else {
				return fmt.Errorf(string(output))
			}
		}

	}
	return nil
}
