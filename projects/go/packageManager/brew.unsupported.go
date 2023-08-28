//go:build !linux || !darwin
// +build !linux !darwin

package packageManager

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/misc/words"
)

// brew - Linux/macOS package manager wrapper function
func brew(pkg DependencyDescriptor) (output string, err error) {
	return words.EmptyString, nil
}
