//go:build !windows
// +build !windows

package packageManager

import (
	"github.com/sam-caldwell/go/v2/projects/go/misc/words"
)

// choco - windows package manager wrapper function
func choco(pkg DependencyDescriptor) (output string, err error) {
	return words.EmptyString, nil
}
