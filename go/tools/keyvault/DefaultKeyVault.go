package main

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"os"
	"path/filepath"
)

// DefaultKeyVault - Return the default keyvault name (e.g. ~/.keyvault.kvt)
//
//	(c) 2023 Sam Caldwell.  MIT License
func DefaultKeyVault() string {
	const fileName = ".keyvault.kvt"
	var home = os.Getenv("HOME")
	if !directory.Exists(home) {
		panic("home directory doesn't exist")
	}
	return filepath.Join(home, fileName)
}
