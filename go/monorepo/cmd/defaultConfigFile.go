package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func defaultConfigFile() string {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		cobra.CheckErr(err)
	}
	return filepath.Join(
		homeDirectory,
		gitRoot,
		defaultRepoName,
		defaultConfigFileName)
}
