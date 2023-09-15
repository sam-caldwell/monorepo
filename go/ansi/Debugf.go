package ansi

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Debugf - Print debug line (with cobra integration for --nocolor flag)
func Debugf(cmd *cobra.Command, format string, msg ...interface{}) {
	Debug(cmd, fmt.Sprintf(format, msg...))
}
