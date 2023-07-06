package packageManager

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
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
