package systemrecon

/*
 * projects/systemrecon/HasExecutable.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This cross-platform function will check to see if a given
 * executable exists on the target host and return an exit code
 * and string answer, where exitCode 0 and answer 'yes' indicate
 * the executable has been found, and where exitCode != 0 or
 * answer 'no' indicates the executable does not exist.
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/v2/misc/words"
	"os/exec"
	"runtime"
	"strings"
)

const (
	/*
	 * Shell code for Windows
	 */
	powershellCode = `if (Get-Command -Name '%s' -ErrorAction SilentlyContinue) { '0' } else { '1' }`
	/*
	 * Shell code for the rest of the world
	 */
	shellCode = "command -v %s >/dev/null 2>&1; echo $?"
)

func HasExecutable(targetCommand string) (exitCode int, answer string) {
	response := words.No
	var cmd *exec.Cmd
	switch goos := runtime.GOOS; goos {
	case "windows":
		cmd = exec.Command("powershell", "-Command", fmt.Sprintf(powershellCode, targetCommand))
	case "darwin", "linux":
		cmd = exec.Command("sh", "-c", fmt.Sprintf(shellCode, targetCommand))
	default:
		return 2, response // unsupported operating system
	}

	out, err := cmd.Output()
	if err == nil && strings.TrimSpace(string(out)) == "0" {
		exitCode = 0
		response = words.Yes
	} else {
		exitCode = 1
		response = words.No
	}
	return exitCode, response
}
