package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// LogTarget - The target to which logs will be written
type LogTarget interface {
	Write(p LogEvent.MessageValue)
	SetLevel(n LogLevel.Value)
	Flush()
}
