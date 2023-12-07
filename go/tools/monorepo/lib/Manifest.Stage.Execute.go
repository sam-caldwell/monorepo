package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os/exec"
	"strings"
	"time"
)

func commandArgs(commandLine []string) (command string, args []string) {
	if len(commandLine[1:]) > 0 {
		args = commandLine[1:]
	}
	return commandLine[0], args
}

func (s *Stage) Execute(className, projectName string, debug bool) error {
	for _, step := range s.Steps {
		if showProjectStatus(step.Enabled, className, projectName, step.Command) {
			continue
		}
		commandLine := strings.Split(step.Command, words.NewLine)
		for _, line := range commandLine {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			command, args := commandArgs(strings.Split(line, words.Space))
			ansi.Cyan().Printf("    line: %s, %v", command, args).LF().Reset()
			output, err := exec.Command(command, args...).CombinedOutput()

			hasError := strings.Contains(strings.ToLower(string(output)), "error")

			if err != nil {
				ansi.Red().Printf("Error (CombinedOutput):%v\n", err).Reset()
				return err
			}

			if debug {
				ansi.White().
					Printf("\tContinueOnError: %v\n", step.ContinueOnError).
					Print("\tOutput:").
					Yellow().LF()
				outputLines := strings.Split(strings.TrimSuffix(string(output), words.NewLine), words.NewLine)
				for _, lineOut := range outputLines {
					ansi.Yellow().Printf("\t%s:", time.Now().Format(time.RFC1123))
					lwrcsLn := strings.ToLower(lineOut)
					if strings.Contains(lwrcsLn, "fail") {
						ansi.Red()
					}
					if strings.Contains(lwrcsLn, "pass") || strings.Contains(lwrcsLn, "ok") {
						ansi.Green()
					}
					ansi.Printf("%s\n", lineOut)
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