package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/systemrecon"
	"github.com/spf13/cobra"
)

func GetCurrentOpsys(cmd *cobra.Command) []string {
	opsys, err := systemrecon.OpSys()
	if err != nil {
		ansi.CheckErr(cmd, err)
	}

	osList := []string{opsys}

	switch opsys {
	case "macos":
		osList = append(osList, "darwin")
	case "ubuntu", "debian", "redhat", "gentoo", "arch":
		osList = append(osList, "linux")
	case "windows7", "windows8", "windows10", "windows11":
		osList = append(osList, "windows")
		//default:
		//	ansi.CheckErr(cmd, fmt.Sprintf("unrecognized opsys (%s)", opsys))
	}
	ansi.Debugf("opsys:%v\n", osList)
	return osList
}
