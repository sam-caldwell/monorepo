package LogTarget

import "github.com/sam-caldwell/monorepo/go/logger/LogLevel"

// StdoutTarget - Send log output of the given format to stdout
//
//	(c) 2023 Sam Caldwell.  MIT License
type StdoutTarget struct {
	level     LogLevel.Value
	rateLimit uint
}
