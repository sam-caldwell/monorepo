package ansi

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
)

// NoColor - Turn off ANSI Color codes. (This must be called for each context when nocolor is respected).
func (c *Color) NoColor(cmd *cobra.Command) *Color {
	c.disabled = cmd.PersistentFlags().Lookup(words.NoColor).Value.String() == words.True
	return c
}

func NoColor(cmd *cobra.Command) *Color {
	c := Color{}
	c.NoColor(cmd)
	return &c //return the given color object
}
