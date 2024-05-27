package LogTarget

import "github.com/sam-caldwell/monorepo/go/logger/LogLevel"

// HttpTarget - Send log output of the given format to Http requests
//
//	(c) 2023 Sam Caldwell.  MIT License
type HttpTarget struct {
	level     LogLevel.Value
	rateLimit uint
}
