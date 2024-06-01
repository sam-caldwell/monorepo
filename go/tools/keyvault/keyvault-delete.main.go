package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
)

// main - keyvault-delete
//
// Perform a secure deletion of a keyvault file.  This means overwriting the existing file byte for byte
// a number of times before deleting the file.
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
}
