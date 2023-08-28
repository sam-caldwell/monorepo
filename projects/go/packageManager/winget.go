//go:build windows
// +build windows

package packageManager

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/runcommand"
	"github.com/sam-caldwell/monorepo/v2/projects/go/systemrecon/opsys"
)

// winget - windows package manager wrapper function
func winget(pkg DependencyDescriptor) (output string, err error) {
	const pattern = "winget install -y %s"
	if pkg.SkipIfDetected {
		if exitCode, _ := systemrecon.HasExecutable(pkg.Name); exitCode == 0 {
			return "dependency exists. skipping", nil
		}
	}
	return runcommand.ShellExecute(fmt.Sprintf(pattern, pkg.Detail))
}
