package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"os/exec"
	"strings"
	"time"
)

func (s *Stage) Execute(debug bool) error {
	for _, step := range s.Steps {
		var args []string
		if strings.TrimSpace(step.Command) == "" {
			continue
		}
		ansi.White().Printf("  [").Bold()
		if step.Enabled {
			ansi.Green().Printf("enabled").Reset().White().Printf("]  step: %s ", step.Command).Reset()
		} else {
			ansi.Yellow().Printf("disabled").Reset().White().Printf("] step: %s ", step.Command).Reset()
			continue
		}
		for _, line := range strings.Split(step.Command, "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			//ansi.Cyan().Printf("line: '%s'", line).LF().Reset()
			parts := strings.Split(line, " ")
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
				for _, lineOut := range strings.Split(strings.TrimSuffix(string(output), "\n"), "\n") {
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
