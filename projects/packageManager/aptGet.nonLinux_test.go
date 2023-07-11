//go:build !linux
// +build !linux

package packageManager

import (
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"testing"
)

func TestAptGetLinux(t *testing.T) {
	data := []DependencyDescriptor{
		{},
		{
			Name:           "test1",
			Installer:      Pkg,
			Detail:         "detail1",
			SkipIfDetected: true,
		},
		{
			Name:           "test2",
			Installer:      Shell,
			Detail:         "detail2",
			SkipIfDetected: true,
		},
		{
			Name:           "test3",
			Installer:      GoGet,
			Detail:         "detail3",
			SkipIfDetected: true,
		},
		{
			Name:           "test4",
			Installer:      GoInstall,
			Detail:         "detail4",
			SkipIfDetected: true,
		},
	}

	for _, descriptor := range data {
		out, err := aptGet(descriptor)
		if out != words.EmptyString {
			t.Fatalf("unexpected output on %s", descriptor.Name)
		}
		if err != nil {
			t.Fatalf("unexpected error on %s", descriptor.Name)
		}
	}
}
