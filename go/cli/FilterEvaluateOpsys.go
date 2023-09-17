package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/spf13/cobra"
	"strings"
)

// EvaluateFilter - Return true if config contains the filter values.
//
// Evaluate the list of expected opsys names against a comma-delimited
// list of opsys and return true if the expected matches any of the
// enabled --opsys <list> values.
func EvaluateFilter(cmd *cobra.Command, configPath string, expectedOpsys []string) bool {
	cliOpsysFlags := strings.Split(
		strings.TrimSpace(cmd.Flags().Lookup(configPath).Value.String()),
		words.Comma)

	ansi.Red().Printf("cliOpsysFlags(%d):%v\n", len(cliOpsysFlags), cliOpsysFlags).Reset()
	ansi.Red().Printf("expectedOpsys:%v\n", expectedOpsys).Reset()
	//for i, lhs := range cliOpsysFlags {
	//	ansi.Green().Printf("%d :: %s\n", i, lhs).Reset()
	//	for j, rhs := range expectedOpsys {
	//		ansi.Green().Printf("[%d.%d] :: %s : %s\n", i, j, lhs, rhs).Reset()
	//		if lhs == rhs {
	//			return true
	//		} else {
	//			ansi.Red().Printf("\tMismatch: %s %s", lhs, rhs).LF().Reset()
	//		}
	//	}
	//}
	return false
}
