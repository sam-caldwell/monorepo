package ansi

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

func CheckErr(cmd *cobra.Command, msg interface{}) {
	if cmd.PersistentFlags().Lookup(words.NoColor).Value.String() != "true" {
		Red()
		defer Reset()
	}
	cobra.CheckErr(msg)
}
