//go:build !linux
// +build !linux

package packageManager

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/misc/words"
)

// aptGet - debian/ubuntu package manager wrapper function
func aptGet(pkg DependencyDescriptor) (output string, err error) {
	return words.EmptyString, nil
}
