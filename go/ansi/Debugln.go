package ansi

import (
	"github.com/spf13/cobra"
)

// Debugln - Print debug line with LF ending (with cobra integration for --debug flag)
func Debugln(cmd *cobra.Command, format string, msg ...interface{}) {
	Debug(cmd, msg)
	LF()
}
