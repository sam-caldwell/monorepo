package LogTarget

import "github.com/sam-caldwell/monorepo/go/logger/LogEvent"

// LogTarget - The target to which logs will be written
type LogTarget interface {
	Write(p LogEvent.LogFormat)
}
