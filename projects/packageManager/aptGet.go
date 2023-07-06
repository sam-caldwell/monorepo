package packageManager

import "github.com/sam-caldwell/go/v2/projects/runcommand"

// aptGet - debian/ubuntu package manager wrapper function
func aptGet(pkg string) (output string, err error) {
	return runcommand.ShellExecute("apt-get install -y --no-install-recommends %s")

}
