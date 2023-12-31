package exit

/*
 * projects/exit/IfVersionRequested.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines IfVersionRequested to
 * detect a version flag in os.Args and
 * show the monorepo version
 *
 * See OpSys.Network.software.Memory.Disk.Cpu.README.md
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
)

// IfVersionRequested - show version and exit program.
func IfVersionRequested() {
	for _, arg := range os.Args {
		if arg == "--version" || arg == "-version" || arg == "-v" {
			fmt.Println(version.Version)
			os.Exit(Success)
		}
	}
}
