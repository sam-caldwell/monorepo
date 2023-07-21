package main

/*
 * get-supported-platforms
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This program lists the operating system and
 * CPU architectures currently supported by the
 * monorepo tools.
 */

import (
	"fmt"
	monorepo "github.com/sam-caldwell/go/v2/projects/__system__"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"os"
	"strings"
)

func main() {
	exit.IfVersionRequested()
	var mode string
	if len(os.Args) >= 2 {
		mode = os.Args[1]
	}
	switch mode {
	case "":
		fmt.Printf("%s\n", strings.Join(monorepo.GetSupportedPlatforms(), ", "))
	case "opsys":
		fmt.Printf("%s\n", strings.Join(monorepo.GetSupportedOpsys(), ", "))
	case "arch":
		fmt.Printf("%s\n", strings.Join(monorepo.GetSupportedArch(), ", "))
	default:
		ansi.Red().Printf("Invalid mode:%s (use opsys, arch or nothing)", mode).LF().Reset()
	}
}
