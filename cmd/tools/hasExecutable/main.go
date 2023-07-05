package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"os"
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

func hasCommand(targetCommand string) (exitCode int, answer string) {
	response := words.No
	fmt.Printf("targetCommand: %s\n", targetCommand)

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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no")
		os.Exit(0)
	}

	targetCommand := strings.TrimSpace(os.Args[1])
	exitCode, answer := hasCommand(targetCommand)
	fmt.Println(answer)

	if len(os.Args) == 3 {
		if strings.TrimSpace(strings.ToLower(os.Args[2])) == "--exitcode" {
			fmt.Println("exit_code:", exitCode)
			os.Exit(exitCode)
		}
	}
}
