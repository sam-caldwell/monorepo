package exit

import (
	"github.com/sam-caldwell/monorepo/projects/ansi"
)

// TerminateOnError - If err is not nil, terminate with GeneralError.
func TerminateOnError(err error) {
	if err != nil {
		/*
		 *
		 */
		ansi.Red().Println(err.Error()).Fatal(GeneralError).Reset()
	}
}
