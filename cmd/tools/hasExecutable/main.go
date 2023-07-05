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
	shellCode      = "%s"
)

func runCommand(shell string, command []string) (exitCode int, response string) {
	response = words.No
	fmt.Printf("shell: %s command: %v\n", shell, command)
	cmd := exec.Command(shell, command...)
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
		return runCommand("powershell", []string{"-Command", fmt.Sprintf(powershellCode, targetCommand)})
	case "darwin":
		return runCommand("command", []string{"-v", fmt.Sprintf(shellCode, targetCommand)})
	case "linux":
		return runCommand("command", []string{"-v", fmt.Sprintf(shellCode, targetCommand)})
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
