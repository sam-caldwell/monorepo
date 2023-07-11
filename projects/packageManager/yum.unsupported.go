//go:build !linux
// +build !linux

package packageManager

import (
	"github.com/sam-caldwell/go/v2/projects/misc/words"
)

// yum - Redhat/centos/fedora package manager wrapper function
func yum(pkg DependencyDescriptor) (output string, err error) {
	return words.EmptyString, nil
}
