//go:build !windows && !linux && !darwin
// +build !windows,!linux,!darwin

package systemrecon

/*
 * getCacheSizes() - Unsupported
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Don't crash, return an error that this is an unsupported operating system.
 *
 * See CpuCache.md
 */
import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/exit/errors"
)

// getCacheSizes - Return a given CPU cache (L1, L2, L3)
func getCacheSizes(level int) (size int, err error) {
	return invalidCacheSz, fmt.Errorf(errors.UnsupportedOpsys)
}
