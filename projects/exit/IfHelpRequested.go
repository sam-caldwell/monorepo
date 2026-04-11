package exit

/*
 * projects/simpleArgs/IfHelpRequested.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This program creates a function which will
 * check os.Args to see if -h or --help was used
 * and if so, the program will print commandUsage
 * and exit.
 *
 * See OpSys.Network.software.Memory.Disk.Cpu.README.md
 */

import (
	"os"
)

// IfHelpRequested - intercept a request for help information
func IfHelpRequested(commandUsage string) {
	for _, arg := range os.Args {
		OnCondition(arg == "-h" || arg == "--help", 0, "", commandUsage)
	}
}
