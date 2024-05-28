package LogTarget

import (
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// SetLevel - Configure the http target with the given log level
func (out *HttpTarget) SetLevel(p logLevel.Value) {
	out.level = p
}
