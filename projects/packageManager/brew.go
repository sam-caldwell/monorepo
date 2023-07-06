package packageManager

import "github.com/sam-caldwell/go/v2/projects/runcommand"

// brew - Linux/macOS package manager wrapper function
func brew(pkg string) (output string, err error) {
	return runcommand.ShellExecute("brew reinstall %s")
}
