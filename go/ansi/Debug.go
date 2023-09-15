package ansi

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

// Debug - Print debug message (with cobra integration for --nocolor flag)
func Debug(cmd *cobra.Command, msg interface{}) {
	if cmd.PersistentFlags().Lookup(words.Debug).Value.String() != words.True {
		return
	}
	if cmd.PersistentFlags().Lookup(words.NoColor).Value.String() != words.True {
		Yellow()
		defer Reset()
	}
	Print(fmt.Sprintf("[Debug]: %v", msg))
}
