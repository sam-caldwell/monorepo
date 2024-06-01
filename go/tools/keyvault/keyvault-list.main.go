package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
)

// main - keyvault-list
//
// Open the keyvault file using passphrase and 2FA token, then list each section
// in the keyvault along with its identifier and payload type.
//
//	(c) 2023 Sam Caldwell.  MIT License
func main() {
	ver := flag.Bool(words.Version, false, "show version string")
	//keyVaultFile := flag.String("keyvault", DefaultKeyVault(), "Specifies the keyvault filename")
	flag.Parse()

	if *ver {
		fmt.Println(version.Version)
		os.Exit(exit.Success)
	}

	//ToDo: add functionality
}
