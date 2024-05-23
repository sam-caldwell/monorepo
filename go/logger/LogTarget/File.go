package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// File - Send log output of the given format to file
type File[LogFormat LogEvent.LogFormat] struct {
	//Todo: define how this will be handled in the file context (need to open file)
}

// SetLevel - Set ANSI color codes when writing log messages.
func (out File[LogFormat]) SetLevel(p logLevel.Value) {
	//Todo: define how this will be handled in the file context
}

// Flush - Flush the log message
func (out File[LogFormat]) Flush() {
	//Todo: define how this will be handled in the file context
}
