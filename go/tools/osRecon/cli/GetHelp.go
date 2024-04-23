package cli

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
	"strings"
)

// GetHelp - Scan os.Args for help commands or --version and exit.
func GetHelp(usage string) {

	for arg := range os.Args {

		s := strings.ToLower(os.Args[arg])

		if (s == "--version") || (s == "version") {
			fmt.Println(version.Version)
			os.Exit(exit.Success)
		}

		if (s == "help") || (s == "-h") || (s == "--help") {
			fmt.Println(usage)
			os.Exit(exit.GeneralError)
		}

	}

}
