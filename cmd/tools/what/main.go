package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/misc/formatting"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	cpu "github.com/sam-caldwell/go/v2/projects/systemrecon/cpu"
	memory "github.com/sam-caldwell/go/v2/projects/systemrecon/memory"
	opsys "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
	"os"
	"strings"
)

const (
	usage = `
what <option> [--printError]

options:
  --printError    : indicates whether error messages are printed to stdout.
  --arch          : Return cpu architecture (e.g. arm64, amd64)
  --cpus          : Return number of CPU cores
  --cpuinfo       : Return CPU specifications
  --meminfo       : Return detailed memory information
  --os            : Return the operating system (e.g. darwin, linux, windows)
  --os-family     : Return an operating system family (e.g. Windows 10)
  --os-version    : Return the operating system version
  --ram           : Return the amount of memory (in KB)
  --software-list : Return a list of software installed on the system
`
)

func main() {
	var output string
	var err error

	intToString := func(fn func() (int, error)) (string, error) {
		raw, err := fn()
		return fmt.Sprintf("%d", raw), err
	}
	mapToString := func(pattern string, fn func() (map[string]string, error)) (output string, err error) {
		data, err := fn()
		if err != nil {
			return output, err
		}
		return misc.FlattenStringMap(data), err
	}

	exit.OnCondition(len(os.Args) < 2, exit.GeneralError, exit.ErrMissingArguments, usage)
	switch command := strings.TrimLeft(strings.ToLower(strings.TrimSpace(os.Args[1])), words.Hyphen); command {

	case "arch":
		output, err = cpu.CpuArch()

	case "cpus":
		output, err = intToString(cpu.CpuCores)

	case "cpuinfo":
		output, err = mapToString(" %10s :%s", cpu.CpuInfo)

	case "meminfo":
		output, err = mapToString(" %10s :%s", memory.MemInfo)

	case "ram":
		output, err = intToString(memory.RamSize)

	case "os":
		output, err = opsys.OpSys()

	case "os-family":
		output, err = opsys.OpSysFamily()

	case "os-version":
		output, err = opsys.OpSysVersion()

	case "help":
		fmt.Printf("Error: %s\n\n%s", err, usage)

	default:
		output = ""
		err = fmt.Errorf(exit.ErrMissingArgWithDetail, command)
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
