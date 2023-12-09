package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
	"os/exec"
	"strings"
)

// showEnvironment - show the key-value environment of the current shell
func showEnvironment(shell *exec.Cmd, expected *[]EnvironmentVariable) {

	if len(*expected) == 0 {

		ansi.Println("\t    OK:No expected environment variables defined")

	} else {

		ansi.Yellow().LF().Printf("Environment Variables:").LF()

		ansi.Printf("\tExpected:\n")
		for n, env := range *expected {
			ansi.Printf("\t    %d %s = %s\n", n, env.Key, env.Value)
		}

		ansi.Println("\t" + strings.Repeat("=", 30))

		ansi.Printf("\tActual:\n")
		discoveredCount := 0
		for n, line := range os.Environ() {
			key, value := splitKeyValue(line)
			for _, expEnv := range *expected {
				if expEnv.Key == key {
					ansi.Printf("\t    %3d %s=%s\n", n, key, value)
					discoveredCount++
				}
			}
		}
		if discoveredCount == len(*expected) {
			ansi.Green().Printf("\t    OK:Expected environment variable count matches\n")
		} else {
			ansi.Red().Printf("\t    Environment is missing parameters\n")
		}
	}
	ansi.LF().Reset()
}

func splitKeyValue(line string) (key, value string) {
	parts := strings.Split(line, words.EqualSign)
	if len(parts) < 2 {
		panic("invalid key=value form of an environment")
	}
	return parts[0], parts[1]
}
