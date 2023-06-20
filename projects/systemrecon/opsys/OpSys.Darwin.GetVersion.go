//go:build darwin
// +build darwin

package systemrecon

import (
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"os/exec"
	"strings"
)

func GetVersion() (version string, err error) {
	out, err := exec.Command("sw_vers", "-productVersion").Output()
	if err != nil {
		return words.EmptyString, err
	}
	version = strings.TrimSpace(string(out))
	return version, nil
}
