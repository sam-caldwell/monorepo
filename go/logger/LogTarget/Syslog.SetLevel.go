package LogTarget

import (
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// SetLevel - Configure the syslog target with the given log level
func (out Syslog) SetLevel(p logLevel.Value) {
	//Noop
}
