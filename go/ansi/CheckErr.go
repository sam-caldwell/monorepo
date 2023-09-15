package ansi

import (
	"github.com/spf13/cobra"
)

func CheckErr(cmd *cobra.Command, msg interface{}) {
	Red()
	defer Reset()
	cobra.CheckErr(msg)
}
