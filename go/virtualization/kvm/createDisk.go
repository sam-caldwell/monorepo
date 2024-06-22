package kvm

import (
	"github.com/sam-caldwell/monorepo/go/fs/file"
)

// CreateDiskImage - creates a disk image using libvirt
//
//	(c) 2023 Sam Caldwell.  MIT License
func CreateDiskImage(path string, size uint64) error {
	return file.CreateRandomImage(path, size)
}
