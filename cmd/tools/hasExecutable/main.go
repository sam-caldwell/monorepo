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
	powershellCode = "if (Get-Command -Name %s -ErrorAction SilentlyContinue) { 'yes' }"
	shellCode      = "%s"
)

func runCommand(shell string, command string, args []string) (exitCode int, response string) {
	response = words.No
	fmt.Printf("shell: %s command: %s args: %v\n", shell, command, args)

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		args = append([]string{"-Command", fmt.Sprintf(powershellCode, command)}, args...)
		cmd = exec.Command(shell, args...)
	} else {
		cmd = exec.Command(shell, append([]string{fmt.Sprintf(shellCode, command)}, args...)...)
	}

	if out, err := cmd.Output(); err == nil {
		fmt.Printf("out: '%s'\n", string(out))
		exitCode = 0
		response = words.Yes
	} else {
		fmt.Printf("out: '%s'\n", string(out))
		exitCode = 1
		response = words.No
	}

	return exitCode, response
}

func hasCommand(targetCommand string) (exitCode int, answer string) {
	switch goos := runtime.GOOS; goos {
	case "windows":
		return runCommand("powershell", targetCommand, nil)
	case "darwin":
		return runCommand("command", "-v", []string{targetCommand})
	case "linux":
		return runCommand("command", "-v", []string{targetCommand})
	}
	//unsupported operating system
	return 2, words.No
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
