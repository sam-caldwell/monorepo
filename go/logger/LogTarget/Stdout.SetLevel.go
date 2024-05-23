package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// SetLevel - Set ANSI color codes when writing log messages.
func (out Stdout[LogFormat]) SetLevel(p logLevel.Value) {
	switch p {
	case logLevel.Debug, logLevel.Info:
		ansi.Blue()
	case logLevel.Warning:
		ansi.Yellow()
	case logLevel.Error, logLevel.Critical, logLevel.Fatal:
		ansi.Red()
	}
}
