package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/fs/file"
)

// File - Send log output of the given format to file
//
//	(c) 2023 Sam Caldwell.  MIT License
type File struct {
	file      file.File
	rateLimit uint
}
