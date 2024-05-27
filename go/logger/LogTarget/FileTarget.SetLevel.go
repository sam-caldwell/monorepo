package LogTarget

import (
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// SetLevel - Configure the file log target with the given log level
func (out FileTarget) SetLevel(p logLevel.Value) {
	//Noop
}
