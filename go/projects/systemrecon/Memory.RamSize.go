//go:build !windows && !linux && !darwin
// +build !windows,!linux,!darwin

package systemrecon

/*
 * SystemInfo ()
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * SystemInfo() for Other systems
 *
 * 	Return a zero value and error for an unsupported system
 */
import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/v2/exit/errors"
)

// RamSize - Return the ram size in KB
func RamSize() (int, error) {
	return 0, fmt.Errorf(errors.UnsupportedOpsys)
}
