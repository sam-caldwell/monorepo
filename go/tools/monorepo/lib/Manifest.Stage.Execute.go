package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/version"
	"os/exec"
	"strings"
	"time"
)

// Execute - Resolve build parameters and execute build steps
func (s *Stage) Execute(rootDir, manifestDir, className, projectName, opsys, arch *string, debug bool) (err error) {
	parameters := map[string]string{
		"${BUILD_ROOT}":    *rootDir,
		"${BUILD_VERSION}": version.Version,
		"${BUILD_OS}":      *opsys,
		"${BUILD_ARCH}":    *arch,
		"${PROJECT_NAME}":  *projectName,
		"${MANIFEST_DIR}":  *manifestDir,
	}

	var hasError bool

	for _, step := range s.Steps {
		//if showProjectStatus(step.Enabled, className, projectName, &step.Command) {
		//	continue
		//}
		commandLine := strings.Split(step.Command, words.NewLine)
		for _, rawLine := range commandLine {
			if rawLine = strings.TrimSpace(rawLine); rawLine == "" {
				continue
			}
			line := resolveParameter(parameters, rawLine)

			command, args := commandArgs(strings.Split(line, words.Space))
			ansi.Cyan().Printf("      └─running: %s, %s", command, strings.Join(args, words.Space)).
				LF().Reset()

			if err = resolveEnvironment(&step.Environment); err != nil {
				ansi.Red().Printf("Error setting environment variables. %v", err).Reset()
				if step.ContinueOnError {
					err = nil
				} else {
					return err
				}
			}

			shell := exec.Command(command, args...)
			setGolangEnvParams(shell, className, opsys, arch)

			if debug {
				//Show the environment variables if debug is set.
				showEnvironment(shell, &step.Environment)
			}

			var output []byte
			output, err = shell.CombinedOutput()
			if err != nil {
				if step.ShowOutput {
					ansi.
						Red().
						Printf("Error (Command shell [output]):%v\n\n%v\n\n", err, string(output)).
						Reset()
				}
				if step.ContinueOnError {
					err = nil
				} else {
					return err
				}
				hasError = strings.Contains(strings.ToLower(string(output)), "error")
			}
			if step.ShowOutput {
				ansi.White().
					Printf("\tContinueOnError: %v\n", step.ContinueOnError).
					Printf("\tShowOutput: %v\n", step.ShowOutput).
					Print("\tOutput:").LF().
					White().Println("Output:")

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
			}
			if hasError {
				if step.ContinueOnError {
					err = nil
				} else {
					return fmt.Errorf("failed: %s", output)
				}
			}
			ansi.Reset()
		}
	}
	return nil
}
