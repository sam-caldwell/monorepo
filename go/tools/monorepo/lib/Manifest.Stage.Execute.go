package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/version"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Execute - Resolve build parameters and execute build steps
func (s *Stage) Execute(rootDir, manifestDir, className, projectName, opsys, arch *string, debug bool) (err error) {
	parameters := map[string]string{
		"${BUILD_ROOT}":    filepath.Join(*rootDir, words.Build),
		"${BUILD_VERSION}": version.Version,
		"${BUILD_OS}":      *opsys,
		"${BUILD_ARCH}":    *arch,
		"${PROJECT_NAME}":  *projectName,
		"${MANIFEST_DIR}":  *manifestDir,
	}

	var hasError bool

	for i, step := range s.Steps {
		commandLine := strings.Split(step.Command, words.NewLine)
		for _, rawLine := range commandLine {
			if rawLine = strings.TrimSpace(rawLine); rawLine == "" {
				continue
			}
			line := resolveParameter(parameters, rawLine)

			command, args := commandArgs(strings.Split(line, words.Space))

			showComment(step.Comment, 1)

			ansi.Cyan().Printf("      ")
			if i == 0 {
				if len(s.Steps) > 0 {
					ansi.Print("└┬")
				} else {
					ansi.Print("└─")
				}
			} else {
				ansi.Print(" ├")
			}

			ansi.Printf("running: %s, %s", command, strings.Join(args, words.Space)).
				LF().Reset()

			if err = resolveEnvironment(&step.Environment); err != nil {
				ansi.Red().Printf("Error setting environment variables. %v", err).Reset()
				if !step.ContinueOnError {
					return fmt.Errorf("error setting environment variables. %v", err)
				}
			}

			var shell *exec.Cmd
			if step.RunAsShell {
				shell = exec.Command("/bin/bash", []string{"-c", line}...)
			} else {
				shell = exec.Command(command, args...)
			}
			setGolangEnvParams(shell, className, opsys, arch)

			if debug {
				//Show the environment variables if debug is set.
				showEnvironment(&step.Environment)
			}

			var output []byte
			output, err = shell.CombinedOutput()
			if err != nil {
				if step.ShowOutput {
					ansi.Red().
						Printf("       ├─Error (Command [output]):%v\n", err).
						Reset()
					outputLines := strings.Split(strings.TrimSuffix(string(output), words.NewLine), words.NewLine)
					for _, lineOut := range outputLines {
						lwrcsLn := strings.ToLower(lineOut)
						if strings.Contains(lwrcsLn, "fail") {
							ansi.Red()
						}
						if strings.Contains(lwrcsLn, "pass") || strings.Contains(lwrcsLn, "ok") {
							ansi.Green()
						}
						if strings.Contains(lwrcsLn, "not implemented") || strings.Contains(lwrcsLn, "disabled") {
							ansi.Magenta()
						}
						ansi.Printf("       ├──%s: ", time.Now().Format(time.RFC1123)).
							Printf("        │  %s\n", lineOut)
					}
				}
				if step.ContinueOnError {
					err = nil
				} else {
					return err
				}
			}
			hasError = strings.Contains(strings.ToLower(string(output)), "error")
			if step.ShowOutput {
				ansi.Cyan().
					Printf(""+
						"       │ContinueOnError : %v\n"+
						"       │ShowOutput      : %v\n"+
						"       │Output:\n", step.ContinueOnError, step.ShowOutput)

				outputLines := strings.Split(strings.TrimSuffix(string(output), words.NewLine), words.NewLine)
				for _, lineOut := range outputLines {
					ansi.
						Yellow().
						Printf("       │   %s: ", time.Now().Format(time.RFC1123))

					lwrcsLn := strings.ToLower(lineOut)
					if strings.Contains(lwrcsLn, "fail") {
						ansi.Red()
					}
					if strings.Contains(lwrcsLn, "pass") || strings.Contains(lwrcsLn, "ok") {
						ansi.Green()
					}
					if strings.Contains(lwrcsLn, "not implemented") || strings.Contains(lwrcsLn, "disabled") {
						ansi.Magenta()
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
