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

// supportedExecutables - A string list
var supportedExecutables = []string{
	words.AptGet, words.Brew, words.Chocolatey, words.Dpkg, words.Rpm, words.Winget, words.Yum, words.Bash,
	words.CommandCom, words.Cshell, words.NodeJs, words.Perl, words.Powershell, words.Python, words.Sh,
	words.Zshell, words.Ftp, words.Scp, words.Ssh,
}

func runCommand(shell, c, args string) (response string) {
	response = words.No
	cmd := exec.Command(shell, c, args)
	_, err := cmd.Output()
	if err == nil {
		response = words.Yes
	}
	return response
}

func hasCommand(targetCommand string) string {
	switch goos := runtime.GOOS; goos {
	case "windows":
		return runCommand(words.Powershell, "-Command", fmt.Sprintf(powershellCode, targetCommand))
	case "darwin":
		fallthrough
	case "linux":
		return runCommand(words.Command, "-v", targetCommand)
	default:
		return words.No
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no")
		os.Exit(0)
	}
	for _, command := range supportedExecutables {
		if command == strings.TrimSpace(os.Args[1]) {
			fmt.Println(hasCommand(command))
			os.Exit(0)
		}
	}
}
