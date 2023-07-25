package packageManager

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
)

// goInstall - Wrapper for 'go get'
func goInstall(pkg DependencyDescriptor) (output string, err error) {
	const pattern = "go install %s"
	return runcommand.ShellExecute(fmt.Sprintf(pattern, pkg.Detail))
}
