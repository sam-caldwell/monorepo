package LogTarget

import (
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// SetLevel - Set ANSI color codes when writing log messages.
func (out File) SetLevel(p logLevel.Value) {
	//Noop
}
