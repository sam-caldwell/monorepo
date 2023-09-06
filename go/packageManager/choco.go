//go:build windows
// +build windows

package packageManager

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/runcommand"
	"github.com/sam-caldwell/monorepo/go/systemrecon"
)

// choco - windows package manager wrapper function
func choco(pkg DependencyDescriptor) (output string, err error) {
	const pattern = "choco install -y %s"
	if pkg.SkipIfDetected {
		if exitCode, _ := systemrecon.HasExecutable(pkg.Name); exitCode == 0 {
			return "dependency exists. skipping", nil
		}
	}
	return runcommand.ShellExecute(fmt.Sprintf(pattern, pkg.Detail))
}
