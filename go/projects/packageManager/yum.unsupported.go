//go:build !linux
// +build !linux

package packageManager

import (
	"github.com/sam-caldwell/monorepo/go/projects/misc/words"
)

// yum - Redhat/centos/fedora package manager wrapper function
func yum(pkg DependencyDescriptor) (output string, err error) {
	return words.EmptyString, nil
}
