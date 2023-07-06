package packageManager

import "github.com/sam-caldwell/go/v2/projects/runcommand"

// choco - windows package manager wrapper function
func choco(pkg string) (output string, err error) {
	return runcommand.ShellExecute("choco install -y %s")
}
