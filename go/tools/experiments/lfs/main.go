package main

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	env "github.com/sam-caldwell/monorepo/go/environment"
	"github.com/sam-caldwell/monorepo/go/misc/size"
	lfs "github.com/sam-caldwell/monorepo/go/tools/experiments/lfs/lib"
	"path/filepath"
)

var partitions = []lfs.Partition{
	{Name: "root.disk", Size: 20 * size.GB, Format: "ext4"},
	{Name: "usr.disk", Size: 20 * size.GB, Format: "ext4"},
	{Name: "var.disk", Size: 20 * size.GB, Format: "ext4"},
	{Name: "var.log.disk", Size: 5 * size.GB, Format: "ext4"},
	{Name: "boot.disk", Size: 500 * size.MB, Format: "ext4"},
	{Name: "boot.efi.disk", Size: 1 * size.MB, Format: "ext4"},
	{Name: "usr.src.disk", Size: 5 * size.GB, Format: "ext4"},
	{Name: "home.disk", Size: 5 * size.GB, Format: "ext4"},
}

func main() {
	homeDir, err := env.RequireString("HOME")
	if err != nil {
		panic(err)
	}
	rootDir := filepath.Join(homeDir, "lfs")
	if err := lfs.CreatePartitions(rootDir, partitions); err != nil {
		panic(err)
	}
	ansi.Green().Println("partitions created").Printf("%v", partitions).LF().Reset()
}
