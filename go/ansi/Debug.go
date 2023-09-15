package ansi

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

// Debug - Print debug message (with cobra integration for --debug flag)
func Debug(cmd *cobra.Command, msg interface{}) {
	if cmd == nil || cmd.PersistentFlags().Lookup(words.Debug).Value.String() != words.True {
		return
	}
	Yellow().Print(fmt.Sprintf("[Debug]: %v", msg)).Reset()
}
