package exit

/*
 * CheckError()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * check if error is nil and if not, print an error message
 * and terminate with GeneralError(1) exit code.
 */

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(GeneralError)
	}
}

func CheckErrorf(err error, fmtString string) {
	if err != nil {
		fmt.Printf(fmtString, err)
		os.Exit(GeneralError)
	}
}
