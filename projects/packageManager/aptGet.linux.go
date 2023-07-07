//go:build linux
// +build linux

package packageManager

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
)

// aptGet - debian/ubuntu package manager wrapper function
func aptGet(pkg DependencyDescriptor) (output string, err error) {
	const pattern = "apt-get install -y --no-install-recommends %s"
	if pkg.SkipIfDetected {
		if exitCode, _ := systemrecon.HasExecutable(pkg.Name); exitCode == 0 {
			return "dependency exists. skipping", nil
		}
	}
	return runcommand.ShellExecute(fmt.Sprintf(pattern, pkg.Detail))
}
