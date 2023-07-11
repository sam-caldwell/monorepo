package argparse

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/wrappers/os"
)

// ExitOnError - Exit on any error
func (arg *Arguments) ExitOnError() *Arguments {
	if arg.HasErrors() {
		fmt.Println("")
		fmt.Println(errArgParseError)
		for _, e := range arg.Errors() {
			fmt.Println(e)
		}
		fmt.Println("")
		fmt.Println(arg.Help())
		os.Exit(exitArgParseError)
	}
	return arg
}
