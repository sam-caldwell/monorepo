//go:build linux
// +build linux

package packageManager

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/v2/runcommand"
	"github.com/sam-caldwell/monorepo/go/projects/v2/systemrecon/opsys"
)

// yum - Redhat/centos/fedora package manager wrapper function
func yum(pkg DependencyDescriptor) (output string, err error) {
	const pattern = "yum install -y %s"
	if pkg.SkipIfDetected {
		if exitCode, _ := systemrecon.HasExecutable(pkg.Name); exitCode == 0 {
			return "dependency exists. skipping", nil
		}
	}
	return runcommand.ShellExecute(fmt.Sprintf(pattern, pkg.Detail))
}
