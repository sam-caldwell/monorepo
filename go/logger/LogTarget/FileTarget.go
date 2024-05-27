package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// FileTarget - Send log output of the given format to file
//
//	(c) 2023 Sam Caldwell.  MIT License
type FileTarget struct {
	file      file.File
	level     LogLevel.Value
	appName   string
	msgId     string
	rateLimit uint
}
