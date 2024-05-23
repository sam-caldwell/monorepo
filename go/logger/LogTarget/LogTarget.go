package LogTarget

import "github.com/sam-caldwell/monorepo/go/logger/LogEvent"

// LogTarget - The target to which logs will be written
type LogTarget interface {
	WriteString(p LogEvent.LogFormat)
}
