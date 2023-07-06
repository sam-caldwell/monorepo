package packageManager

import "github.com/sam-caldwell/go/v2/projects/runcommand"

// winget - windows package manager wrapper function
func winget(pkg string) (output string, err error) {
	return runcommand.ShellExecute("winget install -y %s")
}
