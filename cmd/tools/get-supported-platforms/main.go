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
	monorepo2 "github.com/sam-caldwell/monorepo/v2/projects/go/__system__"
	"github.com/sam-caldwell/monorepo/v2/projects/go/ansi"
	"github.com/sam-caldwell/monorepo/v2/projects/go/exit"
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
		fmt.Printf("%s\n", strings.Join(monorepo2.GetSupportedPlatforms(), ", "))
	case "opsys":
		fmt.Printf("%s\n", strings.Join(monorepo2.GetSupportedOpsys(), ", "))
	case "arch":
		fmt.Printf("%s\n", strings.Join(monorepo2.GetSupportedArch(), ", "))
	default:
		ansi.Red().Printf("Invalid mode:%s (use opsys, arch or nothing)", mode).LF().Reset()
	}
}
