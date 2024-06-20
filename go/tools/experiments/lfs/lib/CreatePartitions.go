package lfs

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/random"
	"path/filepath"
	"time"
)

type Partition struct {
	Name   string
	Size   uint64
	Format string
}

// CreatePartitions - Create a partition loop file
func CreatePartitions(baseDir string, partitions []Partition) (err error) {
	count := 0
	for _, part := range partitions {
		ansi.Blue().Printf("Creating partition %s (size: %d)\n", part.Name, part.Size).Reset()
		go func(n, f string, s uint64) {
			count++
			fileName := filepath.Join(baseDir, fmt.Sprintf("%s", n))
			if err = random.GenerateRandomFile(&fileName, s); err == nil {
				ansi.Blue().Printf("...created (%s) err: %v\n", n, err).Reset()
			} else {
				ansi.Red().Printf("...create failed(%s),%v", n, err).Reset()
				return
			}
			err = FormatPartition(fileName, f)
			if err == nil {
				ansi.Blue().Printf("...formatted(%s)", n).Reset()
			} else {
				ansi.Red().Printf("...format failed(%s),%v", n, err).Reset()
			}
			err = SetPermissions(&fileName)
			count--
		}(part.Name, part.Format, part.Size)
		if err != nil {
			break
		}
	}
	if err != nil {
		ansi.Red().Printf("Error state: %v", err).LF().Reset()
		return err
	}
	for count > 0 {
		ansi.Blue().Printf("Waiting (%d)", count).LF().Reset()
		time.Sleep(1 * time.Second)
	}
	return err
}
