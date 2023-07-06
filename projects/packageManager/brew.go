package packageManager

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
)

// brew - Linux/macOS package manager wrapper function
func brew(pkg DependencyDescriptor) (output string, err error) {
	const pattern = "brew install %s"
	if pkg.SkipIfDetected {
		if exitCode, _ := systemrecon.HasExecutable(pkg.Name); exitCode == 0 {
			return "dependency exists. skipping", nil
		}
	}
	return runcommand.ShellExecute(fmt.Sprintf(pattern, pkg.Detail))
}
