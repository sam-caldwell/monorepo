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

func runCommand(shell, c, args string) (exitCode int, response string) {
	response = words.No
	cmd := exec.Command(shell, c, args)
	_, err := cmd.Output()

	if err == nil {
		exitCode = 0
		response = words.Yes
	} else {
		fmt.Printf("Error: %v\n", err)
		exitCode = 1
		response = words.No
	}
	return exitCode, response
}

func hasCommand(targetCommand string) (exitCode int, answer string) {
	switch goos := runtime.GOOS; goos {
	case "windows":
		return runCommand(words.Powershell, "-Command", fmt.Sprintf(powershellCode, targetCommand))
	case "darwin":
		fallthrough
	case "linux":
		return runCommand(words.Command, "-v", targetCommand)
	default:
		//unsupported operating system
		return 2, words.No
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no")
		os.Exit(0)
	}
	exitCode, answer := hasCommand(strings.TrimSpace(os.Args[1]))
	fmt.Println(answer)
	os.Exit(exitCode)
}
