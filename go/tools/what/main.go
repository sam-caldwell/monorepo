package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/convert"
	exit2 "github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/keyvalue"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/systemrecon"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
	"strings"
)

const (
	usage = `
what <option> [--printError]

options:
  --help          : display usage information
  --version       : display the current version
  --printError    : indicates whether error messages are printed to stdout.

  --cpus          : Return number of CPU cores
  --cpu-arch      : Return cpu architecture (e.g. arm64, amd64)
  --cpu-cache     : Return the L1,L2,L3 CPU cache size in KB (e.g. 32:128:256 for 32 L1, 128 L2 and 256 L3)
  --cpu-info      : Return CPU specifications
  --os            : Return the operating system (e.g. darwin, linux, windows)
  --os-family     : Return an operating system family (e.g. Windows 10)
  --os-version    : Return the operating system version
  --ram           : Return the amount of memory (in KB)qq
  --software-list : Return a list of software installed on the system
`
)

func main() {
	var output string
	var err error

	exit2.IfVersionRequested()
	exit2.OnCondition(len(os.Args) < 2, exit2.GeneralError, errors.MissingArguments, usage)

	switch command := strings.TrimLeft(strings.ToLower(strings.TrimSpace(os.Args[1])), words.Hyphen); command {
	/*
	 * CPU-related commands
	 */
	case "cpus":
		output, err = convert.IntToStringFuncWrapper(systemrecon.CpuCores)

	case "cpu-arch":
		output, err = systemrecon.CpuArch()

	case "cpu-cache":
		output, err = systemrecon.CpuCache()

	case "cpu-info":
		output, err = keyvalue.Interceptor[string, any](systemrecon.CpuInfo)
	/*
	 * operating system stuff
	 */
	case "os":
		output, err = systemrecon.OpSys()

	case "os-family":
		output, err = systemrecon.OpSysFamily()

	case "os-version":
		output, err = systemrecon.OpSysVersion()
	/*
	 * Memory related
	 */
	case "ram":
		output, err = convert.IntToStringFuncWrapper(systemrecon.RamSize)
	/*
	 * General user stuff (help, version, etc)
	 */
	case "help":
		fmt.Printf("\n%s", usage)

	case "version":
		fmt.Printf("%s", version.Version)

	default:
		output = ""
		err = fmt.Errorf(errors.MissingArguments+errors.Details, command)
	}

	if err != nil {
		if (len(os.Args) == 3) && (os.Args[2] == "--printError") {
			fmt.Println(err)
		}
		os.Exit(exit2.GeneralError)
	}

	fmt.Println(output)

	os.Exit(0)

}
