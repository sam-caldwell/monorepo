//go:build !windows
// +build !windows

package packageManager

import (
	"github.com/sam-caldwell/monorepo/go/projects/misc/words"
)

// winget - windows package manager wrapper function
func winget(pkg DependencyDescriptor) (output string, err error) {
	return words.EmptyString, nil
}
