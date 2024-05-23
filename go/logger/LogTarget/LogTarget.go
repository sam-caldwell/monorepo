package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// LogTarget - The target to which logs will be written
type LogTarget interface {
	Init() //ToDo: define how we will initialize each log mode (stdout, file, http, syslog)
	Write(p LogEvent.LogFormat)
	SetLevel(n logLevel.Value)
	Flush()
}
