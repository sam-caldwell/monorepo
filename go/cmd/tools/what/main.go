package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/convert"
	"github.com/sam-caldwell/monorepo/v2/projects/go/exit"
	"github.com/sam-caldwell/monorepo/v2/projects/go/exit/errors"
	"github.com/sam-caldwell/monorepo/v2/projects/go/keyvalue"
	"github.com/sam-caldwell/monorepo/v2/projects/go/misc/words"
	systemrecon2 "github.com/sam-caldwell/monorepo/v2/projects/go/systemrecon/cpu"
	memory "github.com/sam-caldwell/monorepo/v2/projects/go/systemrecon/memory"
	"github.com/sam-caldwell/monorepo/v2/projects/go/systemrecon/opsys"
	"github.com/sam-caldwell/monorepo/v2/projects/go/version"
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

	exit.IfVersionRequested()
	exit.OnCondition(len(os.Args) < 2, exit.GeneralError, errors.MissingArguments, usage)

	switch command := strings.TrimLeft(strings.ToLower(strings.TrimSpace(os.Args[1])), words.Hyphen); command {
	/*
	 * CPU-related commands
	 */
	case "cpus":
		output, err = convert.IntToStringFuncWrapper(systemrecon2.CpuCores)

	case "cpu-arch":
		output, err = systemrecon2.CpuArch()

	case "cpu-cache":
		output, err = systemrecon2.CpuCache()

	case "cpu-info":
		output, err = keyvalue.Interceptor(systemrecon2.CpuInfo)
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
		output, err = convert.IntToStringFuncWrapper(memory.RamSize)
	/*
	 * General user stuff (help, version, etc)
	 */
	case "help":
		fmt.Printf("\n%s", usage)

	case "version":
		fmt.Printf("%s", version.Version)

	default:
		output = ""
		err = fmt.Errorf(errors.MissingArgWithDetail, command)
	}

	if err != nil {
		if (len(os.Args) == 3) && (os.Args[2] == "--printError") {
			fmt.Println(err)
		}
		os.Exit(exit.GeneralError)
	}

	fmt.Println(output)

	os.Exit(0)

}
