//go:build darwin
// +build darwin

package systemrecon

/*
 * RamSize()
 *
 */
import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
)

// RamSize - Return the ram size in KB
func RamSize() (int, error) {
	//ToDo: return RamSize in KB
	return 0, fmt.Errorf(errors.UnsupportedOpsys)
}
