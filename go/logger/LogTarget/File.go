package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// File - Send log output of the given format to file
type File struct {
	//Todo: define how this will be handled in the file context (need to open file)
}

// SetLevel - Set ANSI color codes when writing log messages.
func (out File) SetLevel(p LogLevel.Value) {
	//Todo: define how this will be handled in the file context
}

// Flush - Flush the log message
func (out File) Flush() {
	//Todo: define how this will be handled in the file context
}
