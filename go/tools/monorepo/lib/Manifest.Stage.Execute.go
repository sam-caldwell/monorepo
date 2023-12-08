package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
	"os/exec"
	"strings"
	"time"
)

// commandArgs - parse the commandline for command and arguments
func commandArgs(commandLine []string) (command string, args []string) {
	if len(commandLine[1:]) > 0 {
		args = commandLine[1:]
	}
	return commandLine[0], args
}

// Execute - Resolve build parameters and execute build steps
func (s *Stage) Execute(rootDir, manifestDir, className, projectName, opsys, arch *string, debug bool) error {
	parameters := map[string]string{
		"${BUILD_ROOT}":    *rootDir,
		"${BUILD_VERSION}": version.Version,
		"${BUILD_OS}":      *opsys,
		"${BUILD_ARCH}":    *arch,
		"${PROJECT_NAME}":  *projectName,
		"${MANIFEST_DIR}":  *manifestDir,
	}

	for _, step := range s.Steps {
		//if showProjectStatus(step.Enabled, className, projectName, &step.Command) {
		//	continue
		//}
		commandLine := strings.Split(step.Command, words.NewLine)
		for _, rawLine := range commandLine {
			rawLine = strings.TrimSpace(rawLine)
			if rawLine == "" {
				continue
			}
			line := rawLine
			for key, value := range parameters {
				line = strings.Replace(line, key, value, -1)
			}
			command, args := commandArgs(strings.Split(line, words.Space))
			ansi.
				Cyan().
				Printf("      └─running: %s, %s", command, strings.Join(args, words.Space)).
				LF().
				Reset()

			shell := exec.Command(command, args...)

			//Go-specific environment setup
			if *className == "go" {
				shell.Env = append(os.Environ(),
					fmt.Sprintf("GOOS=%s", *opsys),
					fmt.Sprintf("GOARCH=%s", *arch),
				)
			}
			for _, env := range step.Environment {
				shell.Env = append(os.Environ(),
					fmt.Sprintf("%s=%s", env.Key, env.Value))
			}

			//Show the environment variables if debug is set.
			if debug {
				ansi.Yellow().LF().Printf("Environment Variables:").LF()
				if len(step.Environment) == 0 {
					ansi.Println("  No environment variables defined")
				} else {
					for n, env := range step.Environment {
						ansi.Printf(" %3d %s\n", n, env)
					}
				}
				ansi.LF().Reset()
			}

			output, err := shell.CombinedOutput()
			if err != nil {
				ansi.
					Red().
					Printf("Error (Command shell [output]):%v\n\n%v\n\n", err, string(output)).
					Reset()
				return err
			}
			hasError := strings.Contains(strings.ToLower(string(output)), "error")

			if debug {
				ansi.White().
					Printf("\tContinueOnError: %v\n", step.ContinueOnError).
					Print("\tOutput:").
					Yellow().LF()

				ansi.White().Println("Output:")
				outputLines := strings.Split(strings.TrimSuffix(string(output), words.NewLine), words.NewLine)
				for _, lineOut := range outputLines {
					ansi.Yellow().Printf("%s:", time.Now().Format(time.RFC1123))
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

				if hasError {
					if !step.ContinueOnError {
						return fmt.Errorf(string(output))
					}
				}
			}
		}
	}
	return nil
}
