package packageManager

import "github.com/sam-caldwell/go/v2/projects/runcommand"

// yum - Redhat/centos/fedora package manager wrapper function
func yum(pkg string) (output string, err error) {
	return runcommand.ShellExecute("yum install -y %s")
}
