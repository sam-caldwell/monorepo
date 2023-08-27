package packageManager

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/runcommand"
)

// goGet - Wrapper for 'go get'
func goGet(pkg DependencyDescriptor) (output string, err error) {
	const pattern = "go get %s"
	return runcommand.ShellExecute(fmt.Sprintf(pattern, pkg.Detail))
}
