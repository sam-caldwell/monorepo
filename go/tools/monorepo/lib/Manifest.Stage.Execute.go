package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os/exec"
	"strings"
	"time"
)

func (s *Stage) Execute(className, projectName string, debug bool) error {
	for _, step := range s.Steps {
		var args []string
		command := strings.TrimSpace(strings.TrimSuffix(step.Command, words.NewLine))
		if (command == "") || showProjectStatus(step.Enabled, className, projectName, step.Command) {
			continue
		}
		for _, line := range strings.Split(step.Command, words.NewLine) {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			parts := strings.Split(line, words.Space)
			command := parts[0]
			if len(parts[1:]) > 0 {
				args = parts[1:]
			}
			ansi.Cyan().Printf("    line: %s, %v", command, args).LF().Reset()
			output, err := exec.Command(command, args...).CombinedOutput()
			if err != nil {
				ansi.Red().Printf("Error (CombinedOutput):%v\n", err).Reset()
				return err
			}

			hasError := strings.Contains(strings.ToLower(string(output)), "error")
			if debug {
				ansi.White().
					Printf("\tContinueOnError: %v\n", step.ContinueOnError).
					Print("\tOutput:").
					Yellow().LF()
				outputLines := strings.Split(strings.TrimSuffix(string(output), words.NewLine), words.NewLine)
				for _, lineOut := range outputLines {
					ansi.Printf("\t%s:%s\n", time.Now().Format(time.RFC1123), lineOut)
				}
				ansi.Reset()
			}
			if hasError {
				if !step.ContinueOnError {
					return fmt.Errorf(string(output))
				}
			}
		}
	}
	return nil
}
