package main

/*
 * This utility program can run on multiple operating systems and answer the question
 * "Is this 'executable' installed?" and it will only return a 'yes' or 'no' response.
 * If any error occurs, it replies 'no'
 */
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
)

func runCommand(shell, command string, args []string) (exitCode int, response string) {
	response = words.No
	arguments := append([]string{command}, args...)
	cmd := exec.Command(shell, arguments...)
	if _, err := cmd.Output(); err == nil {
		exitCode = 0
		response = words.Yes
	} else {
		exitCode = 1
		response = words.No
	}
	return exitCode, response
}

func hasCommand(targetCommand string) (exitCode int, answer string) {
	var shell string
	var commandString string
	var arguments []string
	switch goos := runtime.GOOS; goos {
	case "windows":
		shell = "powershell"
		commandString = "-Command"
		arguments = []string{fmt.Sprintf(powershellCode, targetCommand)}
	case "darwin":
		fallthrough
	case "linux":
		shell = "/bin/bash"
		commandString = "command"
		arguments = []string{"-v", targetCommand}
	default:
		//unsupported operating system
		return 2, words.No
	}
	return runCommand(shell, commandString, arguments)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no")
		os.Exit(0)
	}
	exitCode, answer := hasCommand(strings.TrimSpace(os.Args[1]))
	fmt.Println(answer)
	if len(os.Args) == 3 {
		if strings.TrimSpace(strings.ToLower(os.Args[2])) == "--exitcode" {
			os.Exit(exitCode)
		}
	}

}
