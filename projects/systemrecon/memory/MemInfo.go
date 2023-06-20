//go:build !windows && !linux && !darwin
// +build !windows,!linux,!darwin

package systemrecon

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit"
)

/*
 * MemInfo (Other)
 *
 * MemInfo() for Other systems
 *
 * 	Return an empty memory data set for an unsupported system
 */

// MemInfo - Return nil,nil because this system is not specifically implemented.
func MemInfo() (map[string]string, error) {
	return nil, fmt.Errorf(exit.ErrUnsupportedOpsys)
}
