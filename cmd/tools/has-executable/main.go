package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	"github.com/sam-caldwell/go/v2/projects/go/systemrecon/opsys"
	"os"
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

func main() {
	exit.IfVersionRequested()
	if len(os.Args) < 2 {
		fmt.Println("no")
		os.Exit(0)
	}

	targetCommand := strings.TrimSpace(os.Args[1])
	exitCode, answer := systemrecon.HasExecutable(targetCommand)
	fmt.Println(answer)
	if len(os.Args) == 3 {
		if strings.TrimSpace(strings.ToLower(os.Args[2])) == "--exitcode" {
			os.Exit(exitCode)
		}
	}
}
