package ansi

import (
	"github.com/sam-caldwell/monorepo/projects/exit/safety"
	"github.com/spf13/cobra"
)

// CheckErr - given a cobra command, print error in red
//
//	(c) 2023 Sam Caldwell.  MIT License
func CheckErr(cmd *cobra.Command, msg interface{}) {
	Red()
	safety.Defer(func() { Reset() })
	cobra.CheckErr(msg)
}
