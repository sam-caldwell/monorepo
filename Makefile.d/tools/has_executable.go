package main

/*
 * This utility program can run on multiple operating systems and answer the question
 * "Is this 'executable' installed?" and it will only return a 'yes' or 'no' response.
 * If any error occurs, it replies 'no'
 */
import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	yes            = "yes"
	no             = "no"
	powershell     = "powershell"
	command        = "command"
	powershellCode = "if (Get-Command -Name %s -ErrorAction SilentlyContinue) { 'yes' }"
)

var knownCommands = []string{
	"brew", "winget", "choco", "apt-get", "dpkg", "yum", "rpm",
}

func runCommand(shell, c, args string) string {
	cmd := exec.Command(shell, c, args)
	_, err := cmd.Output()
	if err != nil {
		return no
	}
	return yes
}
func hasCommand(targetCommand string) string {
	switch goos := runtime.GOOS; goos {
	case "windows":
		return runCommand(powershell, "-Command", fmt.Sprintf(powershellCode, targetCommand))
	case "darwin":
		fallthrough
	case "linux":
		return runCommand(command, "-v", targetCommand)
	default:
		return no
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no")
		os.Exit(0)
	}
	for _, command := range knownCommands {
		if command == strings.TrimSpace(os.Args[1]) {
			fmt.Println(hasCommand(command))
			os.Exit(0)
		}
	}
}
