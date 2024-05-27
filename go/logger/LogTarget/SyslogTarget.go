package LogTarget

import "github.com/sam-caldwell/monorepo/go/logger/LogLevel"

// SyslogTarget - Send log output of the given format to Syslog
//
//	(c) 2023 Sam Caldwell.  MIT License
type SyslogTarget struct {
	level     LogLevel.Value
	appName   string
	msgId     string
	rateLimit uint
}
