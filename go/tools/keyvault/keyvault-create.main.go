package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
)

// main - keyvault-create
//
// Create a new keyvault file with its associated passphrase and 2FA hardware token.
// The result of this will be an encrypted file with a single header section and a
// configuration section.
//
//	(c) 2023 Sam Caldwell.  MIT License
func main() {
	ver := flag.Bool(words.Version, false, "show version string")
	keyVaultFile := flag.String("keyvault", DefaultKeyVault(), "Specifies the keyvault filename")
	flag.Parse()

	if *ver {
		fmt.Println(version.Version)
		os.Exit(exit.Success)
	}

	kv, err := keyvault.NewVault(keyVaultFile).
		Unlock().
		AddConfig(keyvault.NewConfiguration()).
		Close()
	if err != nil {
		ansi.Red().Printf("failed to create new keyvault.  Error: %v", err).LF().Reset().Fatal(exit.GeneralError)
	}
}
