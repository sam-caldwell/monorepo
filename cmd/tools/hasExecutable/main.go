package main

import (
	"fmt"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
	"os"
	"strings"
)

func main() {
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
